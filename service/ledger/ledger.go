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
	}
	// 创建分类
	if err := db.Create(&data).Error; err != nil {
		return ledger, code.ErrCreateLedgerCategory, err
	}
	//关联账本和分类
	var ledgerCategories []ledger2.LedgerCategory
	for _, k := range data {
		ledgerCategory := ledger2.LedgerCategory{
			LedgerID:   ledger.ID,
			CategoryID: k.ID,
		}
		ledgerCategories = append(ledgerCategories, ledgerCategory)
	}
	if err := db.Create(&ledgerCategories).Error; err != nil {
		return ledger, code.ErrCreateLedgerCategoryRelation, err
	}
	return ledger, code.SUCCESS, nil
}

//GetLedgerList 获取账本
func (*LedgersService) GetLedgerList(userID string) ([]ledger2.Ledger, int, error) {
	//查出自己创建的账本
	var ledgers []ledger2.Ledger
	db := global.DB
	if err := db.Preload("Creator").Where("creator_id = ?", userID).Find(&ledgers).Error; err != nil {
		return ledgers, code.ErrGetLedgerList, err
	}
	//查出自己协作的账本
	var ledgerArr []ledger2.Ledger
	if err := db.Table("ledgers").Preload("Creator").Select("ledgers.*").Joins("left join ledger_users on ledgers.id = ledger_users.ledger_id").Where("ledger_users.user_id = ?", userID).Find(&ledgerArr).Error; err != nil {
		return ledgers, code.ErrGetLedgerList, err
	}
	ledgers = append(ledgers, ledgerArr...)
	return ledgers, code.SUCCESS, nil
}

func (*LedgersService) DeleteLedger(id string) (int, error) {
	db := global.DB
	//先查找账本是否存在
	var ledger ledger2.Ledger
	if err := db.Where("id = ?", id).First(&ledger).Error; err != nil {
		return code.ErrorLedgerNotExist, err
	}
	//先查出关联分类
	var ledgerCategories []ledger2.LedgerCategory
	if err := db.Where("ledger_id = ?", id).Find(&ledgerCategories).Error; err != nil {
		return code.ErrorGetLedgerCategoryRelationList, err
	}
	//删除对应的分类
	for _, k := range ledgerCategories {
		if err := db.Where("id = ?", k.CategoryID).Unscoped().Delete(&ledger2.CategoryLedger{}).Error; err != nil {
			return code.ErrorDeleteLedgerCategory, err
		}
	}
	//删除账本和分类的关联
	if err := db.Where("ledger_id = ?", id).Delete(&ledger2.LedgerCategory{}).Error; err != nil {
		return code.ErrorDeleteLedgerCategoryRelation, err
	}
	if err := db.Where("id = ?", id).Delete(&ledger2.Ledger{}).Error; err != nil {
		return code.ErrorDeleteLedger, err
	}
	return code.SUCCESS, nil
}
