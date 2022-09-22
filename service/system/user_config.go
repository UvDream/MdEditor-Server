package system

import (
	"server/code"
	"server/global"
	"server/models/system"
)

func (*SysUserService) SetUserConfig(userConfig system.UserConfig, userID string) (config system.UserConfig, cd int, err error) {
	//创建用户配置
	if err := global.DB.Create(&userConfig).Error; err != nil {
		return userConfig, code.ErrorSetUserConfig, err
	}
	var user system.User
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return userConfig, code.ErrorFindUser, err
	}
	user.UserConfigID = userConfig.ID
	if err := global.DB.Save(&user).Error; err != nil {
		return userConfig, code.ErrorSetUserConfig, err
	}
	return userConfig, code.SUCCESS, nil
}

func (*SysUserService) EditorUserConfig(userConfig system.UserConfig) (config system.UserConfig, cd int, err error) {
	if err := global.DB.Save(&userConfig).Error; err != nil {
		return userConfig, code.ErrorSetUserConfig, err
	}
	return userConfig, code.SUCCESS, nil
}
func (*SysUserService) GetUserConfig(id string) (config system.UserConfig, cd int, err error) {
	var user system.User
	//查询用户配置
	if err := global.DB.Preload("UserConfig").Where("id = ?", id).First(&user).Error; err != nil {
		return user.UserConfig, code.ErrorFindUserConfig, err
	}
	return user.UserConfig, code.SUCCESS, nil
}
