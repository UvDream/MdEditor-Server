package ledger

import (
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/ledger"
)

func (*LedgersService) GetLedgerCategoryList(id string, types string) ([]ledger.CategoryLedger, int, error) {
	db := global.DB
	var categories []ledger.CategoryLedger
	if err := db.Model(&ledger.CategoryLedger{}).Preload(clause.Associations).Where("ledger_id = ?", id).Where("type = ?", types).Find(&categories).Error; err != nil {
		return []ledger.CategoryLedger{}, code.ErrorGetLedger, err
	}
	//组装分类为树形结构
	data := getCategoryTree(categories)
	return data, code.SUCCESS, nil
}
func getCategoryTree(data []ledger.CategoryLedger) []ledger.CategoryLedger {
	var result []ledger.CategoryLedger
	for _, v := range data {
		if v.ParentID == "" {
			v.Children = getChildren(v.ID, data)
			result = append(result, v)
		}
	}
	return result
}
func getChildren(id string, data []ledger.CategoryLedger) []ledger.CategoryLedger {
	var result []ledger.CategoryLedger
	for _, v := range data {
		if v.ParentID == id {
			v.Children = getChildren(v.ID, data)
			result = append(result, v)
		}
	}
	return result
}

func (*LedgersService) CreateLedgerCategory(category ledger.CategoryLedger) (ledger.CategoryLedger, int, error) {
	db := global.DB
	//查询是否重复了
	var count int64
	if err := db.Model(&ledger.CategoryLedger{}).Where("ledger_id = ?", category.LedgerID).Where("name = ?", category.Name).Where("parent_id = ?", category.ParentID).Count(&count).Error; err != nil {
		return ledger.CategoryLedger{}, code.ErrorCreateCategory, err
	}
	if count > 0 {
		return ledger.CategoryLedger{}, code.ErrCategoryExist, nil
	}
	if err := db.Create(&category).Error; err != nil {
		return ledger.CategoryLedger{}, code.ErrorCreateCategory, err
	}
	return category, code.SUCCESS, nil
}

func (*LedgersService) UpdateLedgerCategory(category ledger.CategoryLedger) (ledger.CategoryLedger, int, error) {
	db := global.DB
	//查询分类是否存在
	var oldCategory ledger.CategoryLedger
	if err := db.Model(&ledger.CategoryLedger{}).Where("id = ?", category.ID).First(&oldCategory).Error; err != nil {
		return ledger.CategoryLedger{}, code.ErrorCategoryNotExist, err
	}
	if err := db.Model(&ledger.CategoryLedger{}).Where("id = ?", category.ID).Updates(&category).Error; err != nil {
		return ledger.CategoryLedger{}, code.ErrorUpdateCategory, err
	}
	return category, code.SUCCESS, nil
}

func (*LedgersService) DeleteLedgerCategory(id string) (int, error) {
	db := global.DB
	//查询分类是否存在
	var oldCategory ledger.CategoryLedger
	if err := db.Model(&ledger.CategoryLedger{}).Where("id = ?", id).First(&oldCategory).Error; err != nil {
		return code.ErrorCategoryNotExist, err
	}
	if err := db.Delete(&oldCategory).Error; err != nil {
		return code.ErrorDeleteCategory, err
	}
	return code.SUCCESS, nil
}
