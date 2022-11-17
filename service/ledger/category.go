package ledger

import (
	"server/code"
	"server/global"
	"server/models/ledger"
)

func (*LedgersService) GetLedgerCategoryList(id string, types string) ([]ledger.CategoryLedger, int, error) {
	db := global.DB
	//查询账本
	var le ledger.Ledger
	if err := db.Model(&ledger.Ledger{}).Preload("Categories", "type = (?)", types).Where("id = ?", id).First(&le).Error; err != nil {
		return []ledger.CategoryLedger{}, code.ErrorGetLedger, err
	}

	//组装分类为树形结构
	data := getCategoryTree(le.Categories)
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
