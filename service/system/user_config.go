package system

import (
	"server/code"
	"server/global"
	"server/models/system"
	"server/utils"
	"time"
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
func (*SysUserService) SetUserInviteCode(userID string) (data string, cd int, err error) {
	var user system.User
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return "", code.ErrorFindUser, err
	}
	user.InviteCode = utils.RandString(6)
	if err := global.DB.Save(&user).Error; err != nil {
		return "", code.ErrorSetUserInviteCode, err
	}
	return user.InviteCode, code.SUCCESS, nil
}
func (*SysUserService) GetUserInviteCode(userID string) (data string, cd int, err error) {
	var user system.User
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return "", code.ErrorFindUser, err
	}
	return user.InviteCode, code.SUCCESS, nil
}
func (*SysUserService) GetUserInviteCodeList(userID string) (data []system.User, cd int, err error) {
	//先根据id查出用户邀请码
	var user system.User
	if err := global.DB.Where("id = ?", userID).Omit("Password").First(&user).Error; err != nil {
		return data, code.ErrorFindUser, err
	}
	//查询邀请码列表
	if err := global.DB.Where("invited_code = ?", user.InviteCode).Omit("Password").Find(&data).Error; err != nil {
		return data, code.ErrorFindUser, err
	}
	return data, code.SUCCESS, nil
}
func (*SysUserService) FillUserInviteCode(userID, inviteCode string) (data system.User, cd int, err error) {
	//先查询邀请码是否存在
	var user system.User
	if err := global.DB.Where("invite_code = ?", inviteCode).First(&user).Error; err != nil {
		return user, code.ErrorFindUser, err
	}
	if user.ID == userID {
		return system.User{}, code.ErrorFindUser, err
	}
	//保存邀请码
	var user2 system.User
	if err := global.DB.Where("id = ?", userID).First(&user2).Preload("Member").Error; err != nil {
		return user2, code.ErrorFindUser, err
	}
	user2.InvitedCode = inviteCode
	if err := global.DB.Save(&user2).Error; err != nil {
		return user2, code.ErrorSetUserInviteCode, err
	}
	//不存在
	if user.MemberID == "" {
		// 将会员时间为今天的时间加一个月
		user.Member.ExpireTime = time.Now().AddDate(0, 1, 0)
		if err := global.DB.Save(&user.Member).Error; err != nil {
			return user, code.ErrorSetUserInviteCode, err
		}
	} else {
		//存在
		user.Member.ExpireTime = user.Member.ExpireTime.AddDate(0, 1, 0)
		if err := global.DB.Save(&user.Member).Error; err != nil {
			return user, code.ErrorSetUserInviteCode, err
		}
	}
	//use2
	if user2.MemberID == "" {
		user2.Member.ExpireTime = time.Now().AddDate(0, 0, 14)
		if err := global.DB.Save(&user2.Member).Error; err != nil {
			return user2, code.ErrorSetUserInviteCode, err
		}
	} else {
		user2.Member.ExpireTime = user2.Member.ExpireTime.AddDate(0, 0, 14)
		if err := global.DB.Save(&user2.Member).Error; err != nil {
			return user2, code.ErrorSetUserInviteCode, err
		}
	}

	return user, code.SUCCESS, nil
}
