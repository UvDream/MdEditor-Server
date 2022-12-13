package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/global"
	"server/models/system"
	"server/utils"
)

//GetUserListService 获取用户列表
func (*SysUserService) GetUserListService(query *system.SysUserRequest) (list interface{}, total int64, mg string, err error) {
	limit := query.PageSize
	offset := query.PageSize * (query.Page - 1)
	var userList []system.User
	db := global.DB.Model(system.User{})
	if query.Username != "" {
		db = db.Where("user_name LIKE ?", "%"+query.Username+"%")
	}
	if query.Nickname != "" {
		db = db.Where("nick_name LIKE ?", "%"+query.Nickname+"%")
	}
	//查询总数
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, "查询用户总数失败", err
	}
	err = db.Omit("password").Limit(limit).Offset(offset).Find(&userList).Error
	if err != nil {
		return nil, 0, "查询用户列表失败", err
	}
	return userList, total, "查询用户列表成功", nil
}

//func (*SysUserService) FindIsMember(userID string) bool {
//	var user system.User
//	db := global.DB
//	if err := db.Where("id = ?", userID).Preload("Member").First(&user).Error; err != nil {
//		return false
//	}
//	//查询会员日期是否过期
//	if user.Member.ExpireTime.Before(time.Now()) {
//		return false
//	}
//	if user.Member.IsMember {
//		return true
//	}
//	return false
//}

func (*SysUserService) GetUserInfoService(userId string) (user system.User, cd int, err error) {
	var u system.User
	db := global.DB
	if err := db.Where("id = ?", userId).Omit("Password").First(&u).Error; err != nil {
		return u, code.ErrorUserNotExist, err
	}
	return u, code.SUCCESS, nil
}

func (*SysUserService) UnbindEmailService(userId string) (cd int, err error) {
	var u system.User
	db := global.DB
	if err := db.Where("id = ?", userId).First(&u).Error; err != nil {
		return code.ErrorUserNotExist, err
	}
	u.Email = ""
	if err := db.Save(&u).Error; err != nil {
		return code.ErrorUpdateUser, err
	}
	return code.SUCCESS, nil
}

func (*SysUserService) BindEmailService(c *gin.Context, userID string, bindEmail system.BindEmail) (cd int, err error) {
	// 首先查询用户是否存在
	var u system.User
	db := global.DB
	if err := db.Where("id = ?", userID).First(&u).Error; err != nil {
		return code.ErrorUserNotExist, err
	}
	// 查询邮箱是否已经被绑定
	if err := db.Where("email = ?", bindEmail.Email).First(&u).Error; err == nil {
		return code.ErrorEmailExist, err
	}
	cd, err = utils.VerifyEmailCodeService(c, bindEmail.Code, bindEmail.UniqueID, bindEmail.Email)
	if err != nil {
		return cd, err
	}
	u.Email = bindEmail.Email
	if err := db.Save(&u).Error; err != nil {
		return code.ErrorUpdateUser, err
	}
	return code.SUCCESS, nil
}
