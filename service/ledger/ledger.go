package ledger

import (
	"server/code"
	"server/global"
	ledger2 "server/models/ledger"
)

type LedgersService struct{}

//CreateLedger 创建账本
func (*LedgersService) CreateLedger(ledger ledger2.Ledger) (ledger2.Ledger, int, error) {
	//TODO 校验会员和非会员账本数量 会员可创建5个账本 非会员可创建1个账本
	db := global.DB
	if err := db.Where("name = ?", ledger.Name).Where("creator_id = ?", ledger.CreatorID).First(&ledger2.Ledger{}).Error; err == nil {
		return ledger, code.ErrLedgerExist, err
	}
	if err := db.Create(&ledger).Error; err != nil {
		return ledger, code.ErrCreateLedger, err
	}
	//创建账本完成后创建一些默认的分类
	data := ledger2.InitLedgerCategory
	for _, k := range data {
		k.UserID = ledger.CreatorID
		k.IconType = "3"
		k.LedgerID = ledger.ID
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
	if err := db.Preload("Creator").Preload("Categories").Where("creator_id = ?", userID).Find(&ledgers).Error; err != nil {
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
