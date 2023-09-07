package system

import (
	"server/code"
	"server/global"
	"server/models/system"
)

func (*SysUserService) MenuListService() (data []system.Permission, cd int, err error) {
	if err := global.DB.Find(&data).Error; err != nil {
		return data, code.ErrorMenuList, err
	}
	return data, code.SUCCESS, nil
}
func (*SysUserService) DeleteMenuService(id string) (cd int, err error) {
	//先查询是否存在
	var menu system.Permission
	if err := global.DB.Where("id = ?", id).First(&menu).Error; err != nil {
		return code.ErrorDeleteMenuFail, err
	}
	if err := global.DB.Where("id = ?", id).Delete(&system.Permission{}).Error; err != nil {
		return code.ErrorDeleteMenuFail, err
	}
	return code.SUCCESS, err
}

func (*SysUserService) AddMenuService(data system.Permission) (cd int, err error) {
	//先查询是否存在,存在则更新
	var menu system.Permission
	db := global.DB
	if err := db.Where("key = ?", data.Key).Where("name = ?", data.Name).First(&menu).Error; err == nil {
		if err := db.Model(&menu).Updates(data).Error; err != nil {
			return code.ErrorAddMenuFail, err
		}
		return code.SUCCESS, err
	}
	if err := db.Create(&data).Error; err != nil {
		return code.ErrorAddMenuFail, err
	}
	return code.SUCCESS, err
}

func (*SysUserService) UserRoleService(query system.UserRoleRequest) (cd int, err error) {
	//先删除用户角色关系
	db := global.DB
	if err := db.Where("user_id = ?", query.UserID).Delete(&system.UserRole{}).Error; err != nil {
		return code.ErrorUserRoleFail, err
	}
	//再添加用户角色关系
	for _, v := range query.RolesID {
		var userRole system.UserRole
		userRole.UserID = query.UserID
		userRole.RoleID = v
		if err := db.Create(&userRole).Error; err != nil {
			return code.ErrorUserRoleFail, err
		}
	}
	return code.SUCCESS, err
}

func (*SysUserService) GetPermissionService(userId string) (data []string, cd int, err error) {
	//先查询用户角色关系
	db := global.DB
	var userRole []system.UserRole
	if err := db.Where("user_id = ?", userId).Find(&userRole).Error; err != nil {
		return data, code.ErrorGetUser, err
	}
	var roles []string
	for _, v := range userRole {
		roles = append(roles, v.RoleID)
	}
	//再查询角色权限关系
	var rolePermission []system.RolePermission
	if err := db.Where("role_id in ?", roles).Find(&rolePermission).Error; err != nil {
		return data, code.ErrorGetUser, err
	}
	var permissions []string
	for _, v := range rolePermission {
		permissions = append(permissions, v.PermissionID)
	}
	//最后查询权限
	var permission []system.Permission
	if err := db.Where("id in ?", permissions).Find(&permission).Error; err != nil {
		return data, code.ErrorGetUser, err
	}
	for _, v := range permission {
		data = append(data, v.Key)
	}
	return data, cd, err
}
