package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/global"
	"server/models/ledger"
	"server/utils"
)

func (*LedgersService) AddBillService(bill ledger.Bill) (ledger.Bill, int, error) {
	db := global.DB
	if err := db.Create(&bill).Error; err != nil {
		return bill, code.ErrCreateBill, err
	}
	return bill, 0, nil
}

func (*LedgersService) DeleteBillService(id string) (int, error) {
	db := global.DB
	//查询是否存在
	if err := db.Where("id = ?", id).First(&ledger.Bill{}).Error; err != nil {
		return code.ErrorGetBill, err
	}
	if err := db.Where("id = ?", id).Delete(&ledger.Bill{}).Error; err != nil {
		return code.ErrorDeleteBill, err
	}
	return code.SUCCESS, nil
}
func (*LedgersService) UpdateBillService(bill ledger.Bill) (ledger.Bill, int, error) {
	db := global.DB
	//TODO 查询是否有权限更新
	//先查询是否存在
	if err := db.Where("id = ?", bill.ID).First(&ledger.Bill{}).Error; err != nil {
		return bill, code.ErrorGetBill, err
	}
	if err := db.Model(&ledger.Bill{}).Where("id = ?", bill.ID).Updates(&bill).Error; err != nil {
		return bill, code.ErrorUpdateBill, err
	}
	return bill, 0, nil
}
func (*LedgersService) GetBillListService(query ledger.BillRequest, userID string, c *gin.Context) (bill ledger.Bill, total int64, cd int, err error) {
	db := global.DB
	//TODO 查询是否有权限获取
	//先查询账本是否存在
	if err := db.Where("id = ?", query.LedgerID).First(&ledger.Ledger{}).Error; err != nil {
		return bill, total, code.ErrorGetLedger, err
	}
	//查询账单
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	//查询总数
	if err := db.Model(&ledger.Bill{}).Where("ledger_id = ?", query.LedgerID).Count(&total).Error; err != nil {
		return bill, total, code.ErrorGetBill, err
	}
	if err := db.Where("ledger_id = ?", query.LedgerID).Scopes(utils.Paginator(c)).Find(&bill).Error; err != nil {
		return bill, total, code.ErrorGetBill, err
	}
	return bill, total, code.SUCCESS, err
}
