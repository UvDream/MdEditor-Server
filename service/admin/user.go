package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/system"
	service "server/service/system"
	"server/utils"
)

type UserService struct{}

// GetUserList 获取用户列表
func (u *UserService) GetUserList(query system.SysUserRequest, c *gin.Context) (userList []system.User, total int64, cd int, err error) {
	db := global.DB
	if query.Username != "" {
		db = db.Where("user_name LIKE ?", "%"+query.Username+"%")
	}
	if query.Nickname != "" {
		db = db.Where("nick_name LIKE ?", "%"+query.Nickname+"%")
	}
	if query.Email != "" {
		db = db.Where("email LIKE ?", "%"+query.Email+"%")
	}
	if query.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+query.Phone+"%")
	}
	if query.Gender != "" {
		db = db.Where("gender LIKE ?", "%"+query.Gender+"%")
	}
	if err = db.Model(&system.User{}).Count(&total).Error; err != nil {
		return nil, 0, code.ErrorUserList, err
	}
	if err = db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&userList).Error; err != nil {
		return nil, 0, code.ErrorUserList, err
	}

	return userList, total, code.SUCCESS, err
}

// AddRole 新增角色
func (u *UserService) AddRole(role system.Role) (err error) {
	//先查询是否存在该角色
	if !errors.Is(global.DB.Where("role_name = ?", role.RoleName).First(&system.Role{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色名")
	}
	if err = global.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

// GetRoleList 角色列表
func (u *UserService) GetRoleList(c *gin.Context) (roleList []system.Role, total int64, cd int, err error) {
	roleName := c.Query("roleName")
	db := global.DB
	if roleName != "" {
		db = db.Where("role_name LIKE ?", "%"+roleName+"%")
	}
	var userService service.SysUserService
	userId := utils.FindUserID(c)
	roles, err := userService.FindUserRoles(userId)
	if err != nil {
		return nil, 0, code.ErrorRoleList, err
	}
	if len(roles) > 0 {
		//查询roles中是否存在admin/root
		for _, v := range roles {
			if v == "admin" || v == "root" {
				return nil, 0, code.ErrorRoleList, err
			}
		}
	}

	//查询角色列表
	if err = db.Model(&system.Role{}).Count(&total).Error; err != nil {
		return nil, 0, code.ErrorRoleList, err
	}
	//查询角色列表
	if err = db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&roleList).Error; err != nil {
		return nil, 0, code.ErrorRoleList, err
	}
	return roleList, total, code.SUCCESS, err
}
