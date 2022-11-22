package ledger

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

func (*LedgersService) GetBillListService(query ledger.BillRequest, userID string, c *gin.Context) (data []ledger.BillChildren, total int64, cd int, err error) {
	db := global.DB
	var bill []ledger.Bill
	//TODO 查询是否有权限获取
	//先查询账本是否存在
	if err := db.Where("id = ?", query.LedgerID).First(&ledger.Ledger{}).Error; err != nil {
		return data, total, code.ErrorGetLedger, err
	}
	//时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	//查询账单
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	//查询总数
	if err := db.Model(&ledger.Bill{}).Where("ledger_id = ?", query.LedgerID).Count(&total).Error; err != nil {
		return data, total, code.ErrorGetBill, err
	}
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Where("ledger_id = ?", query.LedgerID).Order("create_time desc").Find(&bill).Error; err != nil {
		return data, total, code.ErrorGetBill, err
	}
	var newData map[string]ledger.BillChildren
	newData = make(map[string]ledger.BillChildren)
	for _, v := range bill {
		time := v.CreateTime.Format("2006-01-02")
		if v.Type == "0" {
			newData[time] = ledger.BillChildren{
				Income:      newData[time].Income + v.Amount,
				Expenditure: newData[time].Expenditure,
				Bill:        append(newData[time].Bill, v),
			}
		} else {
			newData[time] = ledger.BillChildren{
				Income:      newData[time].Income,
				Expenditure: newData[time].Expenditure + v.Amount,
				Bill:        append(newData[time].Bill, v),
			}
		}
	}
	for k, v := range newData {
		v.Time = k
		data = append(data, v)
	}
	return data, total, code.SUCCESS, err
}

func (*LedgersService) GetBillDetailService(id string) (bill ledger.Bill, cd int, err error) {
	db := global.DB
	//查询是否存在
	if err := db.Where("id = ?", id).First(&ledger.Bill{}).Error; err != nil {
		return bill, code.ErrorGetBill, err
	}
	if err := db.Where("id = ?", id).Preload(clause.Associations).First(&bill).Error; err != nil {
		return bill, code.ErrorGetBill, err
	}
	return bill, code.SUCCESS, nil
}
