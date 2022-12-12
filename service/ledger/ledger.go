package ledger

import (
	"errors"
	"gorm.io/gorm"
	"server/code"
	"server/global"
	ledger2 "server/models/ledger"
	"server/models/system"
	"server/utils"
	"time"
)

type LedgersService struct{}

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

//CreateLedger 创建账本
func (*LedgersService) CreateLedger(ledger ledger2.Ledger) (ledger2.Ledger, int, error) {
	//校验会员和非会员账本数量 会员可创建5个账本 非会员可创建1个账本
	db := global.DB
	isMember := FindIsMember(ledger.CreatorID, db)
	var count int64
	if err := db.Where("creator_id = ?", ledger.CreatorID).Find(&ledger2.Ledger{}).Count(&count).Error; err != nil {
		return ledger, code.ErrCreateLedger, err
	}
	//不是会员
	if !isMember {
		if count > 1 {
			return ledger, code.ErrorCreateLedgerNotMember, errors.New("非会员只能创建一个账本")
		}
	}
	//是会员
	if count > 5 {
		return ledger, code.ErrorCreateLedgerMember, errors.New("会员只能创建五个账本")
	}
	if err := db.Where("name = ?", ledger.Name).Where("creator_id = ?", ledger.CreatorID).First(&ledger2.Ledger{}).Error; err == nil {
		return ledger, code.ErrLedgerExist, err
	}
	ledger.ShareCodeTime = time.Now()
	if err := db.Create(&ledger).Error; err != nil {
		return ledger, code.ErrCreateLedger, err
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
		return ledger, code.ErrCreateLedgerCategory, err
	}
	return ledger, code.SUCCESS, nil
}

//GetLedgerList 获取账本
func (*LedgersService) GetLedgerList(userID string) ([]ledger2.Ledger, int, error) {
	//查出自己创建的账本
	var ledgers []ledger2.Ledger
	db := global.DB
	if err := db.Preload("Creator").Preload("Categories").Preload("User").Where("creator_id = ?", userID).Find(&ledgers).Error; err != nil {
		return ledgers, code.ErrGetLedgerList, err
	}
	//查出自己协作的账本
	var ledgerArr []ledger2.Ledger
	if err := db.Table("ledgers").Preload("Creator").Select("ledgers.*").Joins("left join ledger_users on ledgers.id = ledger_users.ledger_id").Where("ledger_users.user_id = ?", userID).Find(&ledgerArr).Error; err != nil {
		return ledgers, code.ErrGetLedgerList, err
	}
	ledgers = append(ledgers, ledgerArr...)
	//去重
	ledgers = removeDuplicateLedger(ledgers)
	//查询账本成员数量
	for i, k := range ledgers {
		ledgers[i].MemberCount = k.MemberCount + findMemberCount(k.ID)
	}
	//ledgers = append(ledgers, ledgerArr...)
	return ledgers, code.SUCCESS, nil
}

//去重
func removeDuplicateLedger(ledgers []ledger2.Ledger) []ledger2.Ledger {
	result := make([]ledger2.Ledger, 0, len(ledgers))
	temp := map[string]struct{}{}
	for _, item := range ledgers {
		l := len(temp)
		temp[item.ID] = struct{}{}
		if len(temp) != l {
			result = append(result, item)
		}
	}
	return result
}
func findMemberCount(id string) (count int64) {
	db := global.DB
	if err := db.Where("ledger_id = ?", id).Find(&ledger2.LedgerUser{}).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (*LedgersService) DeleteLedger(id string) (int, error) {
	db := global.DB
	//先查找账本是否存在
	var ledger ledger2.Ledger
	if err := db.Where("id = ?", id).First(&ledger).Error; err != nil {
		return code.ErrorLedgerNotExist, err
	}
	//根据账本id删除分类
	if err := db.Where("ledger_id = ?", id).Delete(&ledger2.CategoryLedger{}).Error; err != nil {
		return code.ErrorDeleteLedgerCategory, err
	}
	if err := db.Where("id = ?", id).Delete(&ledger2.Ledger{}).Error; err != nil {
		return code.ErrorDeleteLedger, err
	}
	return code.SUCCESS, nil
}

func (*LedgersService) GetLedgerDetail(id string) (ledger ledger2.Ledger, cd int, err error) {
	db := global.DB
	if err := db.Preload("Creator").Preload("Categories").Where("id = ?", id).First(&ledger).Error; err != nil {
		return ledger, code.ErrorLedgerNotExist, err
	}
	return ledger, code.SUCCESS, nil
}

func (*LedgersService) ShareLedger(id string, userID string) (string, int, error) {
	db := global.DB
	//判断账本是否存在
	var ledger ledger2.Ledger
	if err := db.Where("id = ?", id).First(&ledger).Error; err != nil {
		return "", code.ErrorLedgerNotExist, err
	}
	//判断是否是账本创建者
	if ledger.CreatorID != userID {
		return "", code.ErrorNotLedgerCreator, errors.New("不是账本创建者")
	}
	//判断共享码是否超过7天
	if ledger.ShareCode != "" && time.Now().Sub(ledger.ShareCodeTime) < 7*24*time.Hour {
		return ledger.ShareCode, code.SUCCESS, nil
	}
	//生成分享码
	shareCode := utils.RandString(6)
	//更新账本分享码以及生成时间
	if err := db.Model(&ledger).Where("id = ?", id).Update("share_code", shareCode).Update("share_code_time", time.Now()).Error; err != nil {
		return "", code.ErrorUpdateLedger, err
	}
	return shareCode, code.SUCCESS, nil
}

func (*LedgersService) JoinLedger(shareCode string, userID string) (int, error) {
	db := global.DB
	//判断账本是否存在
	var ledger ledger2.Ledger
	if err := db.Where("share_code = ?", shareCode).First(&ledger).Error; err != nil {
		return code.ErrorLedgerNotExist, err
	}
	//判断是否已经加入
	var ledgerUser ledger2.LedgerUser
	if err := db.Where("ledger_id = ? AND user_id = ?", ledger.ID, userID).First(&ledgerUser).Error; err != nil {
		//没有加入则加入
		ledgerUser.LedgerID = ledger.ID
		ledgerUser.UserID = userID
		ledgerUser.Permission = "0"
		if err := db.Create(&ledgerUser).Error; err != nil {
			return code.ErrorJoinLedger, err
		}
	}
	return code.SUCCESS, nil
}

func (*LedgersService) InviteLedger(email string, ledgerID string, userID string) (int, error) {
	db := global.DB
	// 查找用户是否存在
	var user system.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return code.ErrorUserNotExist, err
	}
	//判断是否已经加入
	var ledgerUser ledger2.LedgerUser
	if err := db.Where("ledger_id = ? AND user_id = ?", ledgerID, user.ID).First(&ledgerUser).Error; nil == err {
		return code.ErrorUserAlreadyJoined, nil
	}
	//判断是否是账本创建者
	var ledger ledger2.Ledger
	if err := db.Where("id = ?", ledgerID).First(&ledger).Error; err != nil {
		return code.ErrorLedgerNotExist, err
	}
	if ledger.CreatorID != userID {
		return code.ErrorNotLedgerCreator, errors.New("不是账本创建者")
	}
	//发送邮件
	//	if err := utils.SendEmail(email, "账本邀请", "您被邀请加入账本"); err != nil {
	//		return code.ErrorSendEmail, err
	//	}
	//创建账本用户
	ledgerUser.LedgerID = ledgerID
	ledgerUser.UserID = user.ID
	if err := db.Create(&ledgerUser).Error; err != nil {
		return code.ErrorJoinLedger, err
	}
	return code.SUCCESS, nil
}
