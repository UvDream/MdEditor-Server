package system

import (
	"server/global"
	"server/models/system"
	"time"
)

// FindIsMember 查询是否是会员
func (*SysUserService) FindIsMember(userID string) (isMember bool) {
	var user system.User
	db := global.DB
	if err := db.Where("id = ?", userID).Preload("Member").First(&user).Error; err != nil {
		return false
	}
	//查询会员日期是否过期
	if user.Member.ExpireTime.Before(time.Now()) {
		return false
	}
	if user.Member.IsMember {
		return true
	}
	return false
}
