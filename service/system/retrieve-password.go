package system

import (
	"server/code"
	"server/global"
	"server/models/system"
)

func (*SysUserService) RetrievePasswordService(dat system.RetrievePasswordRequest) (num int, err error) {
	db := global.DB
	//查询
	var user system.User
	if err := db.Where("user_name = ? AND nick_name = ? AND phone = ? AND email = ?", dat.UserName, dat.NickName, dat.Phone, dat.Email).First(&user).Error; err != nil {
		return code.ErrorUpdatePasswordMissingParam, err
	}
	//更新
	user.Password = dat.Password
	if err := db.Save(&user).Error; err != nil {
		return code.ErrorUpdatePassword, err
	}
	//if err := db.Model(&user).Update("password", dat.Password).Error; err != nil {
	//	return code.ErrorUpdatePassword, err
	//}
	return code.SUCCESS, nil
}
