package system

import (
	"errors"
	"gorm.io/gorm"
	code2 "server/code"
	"server/global"
	ledger2 "server/models/ledger"
	"server/models/system"
	"time"
)

// RegisterService 注册用户
func (*SysUserService) RegisterService(opts system.User) (user system.User, code int, err error) {
	db := global.DB
	//首先查询账号是否存在
	if err := db.Where("user_name = ?", opts.UserName).First(&user).Error; err == nil {
		return opts, code2.ErrorUserExist, err
	}
	//查询昵称是否存在
	if err := db.Where("nick_name = ?", opts.NickName).First(&user).Error; err == nil {
		return opts, code2.ErrorUserExist, err
	}
	//查询邮箱是否存在
	if opts.Email != "" {
		if err := db.Where("email = ?", opts.Email).First(&user).Error; err == nil {
			return opts, code2.ErrorUserExistEmail, err
		}
	}

	//查询手机号是否存在
	if opts.Phone != "" {
		if err := db.Where("phone = ?", opts.Phone).First(&user).Error; err == nil {
			return opts, code2.ErrorUserExistPhone, err
		}
	}
	//查询性别
	if opts.Gender == "" {
		opts.Avatar = "https://pic.imgdb.cn/item/64c0cc451ddac507ccd49532.png"
	}
	//男
	if opts.Gender == "1" {
		opts.Avatar = "https://pic.imgdb.cn/item/64c0cc451ddac507ccd49532.png"
	}
	//女
	if opts.Gender == "2" {
		opts.Avatar = "https://pic.imgdb.cn/item/64c0cc4a1ddac507ccd49d1c.png"
	}
	if err := db.Create(&opts).Error; err != nil {
		return opts, code2.ErrorCreateUser, err
	}
	//邀请码是否存在
	if opts.InvitedCode != "" {
		var sysService = SysUserService{}
		user, cd, err := sysService.FillUserInviteCode(opts.ID, opts.InvitedCode)
		if err != nil {
			return user, cd, err
		}
	}
	//饭米粒记账用户注册
	if opts.Source == "account" {
		var ledger = ledger2.Ledger{
			Name:        "默认账本",
			CreatorID:   opts.ID,
			Description: "默认账本",
		}
		_, cd, err := CreateLedger(ledger)
		if err != nil {
			return opts, cd, err
		}
	}
	return opts, code2.SUCCESS, err
}
func CreateLedger(ledger ledger2.Ledger) (ledger2.Ledger, int, error) {
	//校验会员和非会员账本数量 会员可创建5个账本 非会员可创建1个账本
	db := global.DB
	isMember := FindIsMember(ledger.CreatorID, db)
	var count int64
	if err := db.Where("creator_id = ?", ledger.CreatorID).Find(&ledger2.Ledger{}).Count(&count).Error; err != nil {
		return ledger, code2.ErrCreateLedger, err
	}
	//不是会员
	if !isMember {
		if count > 1 {
			return ledger, code2.ErrorCreateLedgerNotMember, errors.New("非会员只能创建一个账本")
		}
	}
	//是会员
	if count > 5 {
		return ledger, code2.ErrorCreateLedgerMember, errors.New("会员只能创建五个账本")
	}
	if err := db.Where("name = ?", ledger.Name).Where("creator_id = ?", ledger.CreatorID).First(&ledger2.Ledger{}).Error; err == nil {
		return ledger, code2.ErrLedgerExist, err
	}
	ledger.ShareCodeTime = time.Now()
	if err := db.Create(&ledger).Error; err != nil {
		return ledger, code2.ErrCreateLedger, err
	}
	//创建账本完成后创建一些默认的分类
	data := ledger2.InitLedgerCategory
	for i, _ := range data {
		data[i].UserID = ledger.CreatorID
		data[i].IconType = "3"
		data[i].LedgerID = ledger.ID
	}
	// 创建分类
	if err := db.Create(&data).Error; err != nil {
		return ledger, code2.ErrCreateLedgerCategory, err
	}
	return ledger, code2.SUCCESS, nil
}
func FindIsMember(userID string, db *gorm.DB) bool {
	var user system.User
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
