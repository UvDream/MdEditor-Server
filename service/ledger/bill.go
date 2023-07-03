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
	return bill, code.SUCCESS, nil
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
	if err := db.Model(&ledger.Bill{}).Where("id = ?", bill.ID).Omit("creator_id").Updates(&bill).Error; err != nil {
		return bill, code.ErrorUpdateBill, err
	}
	return bill, code.SUCCESS, nil
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
	//查询金额
	if query.Amount != "" {
		db = db.Where("amount = ?", query.Amount)
	}
	//查询总数
	if err := db.Model(&ledger.Bill{}).Where("ledger_id = ?", query.LedgerID).Count(&total).Error; err != nil {
		return data, total, code.ErrorGetBill, err
	}
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Where("ledger_id = ?", query.LedgerID).Order("create_time desc").Find(&bill).Error; err != nil {
		return data, total, code.ErrorGetBill, err
	}
	for i, v := range bill {
		date := v.CreateTime.Format("2006-01-02")
		if v.Category.ParentID == "" {
			bill[i].CategoryName = v.Category.Name
		} else {
			dp := global.DB
			//	根据id查询父级
			var parent ledger.CategoryLedger
			if err := dp.Where("id = ?", v.Category.ParentID).First(&parent).Error; err != nil {
				return data, total, code.ErrorGetCategory, err
			}
			bill[i].CategoryName = parent.Name + "-" + v.Category.Name
		}
		exit, index := isExit(data, date)
		if exit {
			data[index].Bill = append(data[index].Bill, bill[i])
			if v.Type == "0" {
				data[index].Expenditure += v.Amount
			} else {
				data[index].Income += v.Amount
			}
		} else {
			if v.Type == "0" {
				data = append(data, ledger.BillChildren{
					Time:        date,
					Expenditure: v.Amount,
					Bill:        []ledger.Bill{bill[i]},
				})
			} else {
				data = append(data, ledger.BillChildren{
					Time:   date,
					Income: v.Amount,
					Bill:   []ledger.Bill{bill[i]},
				})
			}
		}
	}

	return data, total, code.SUCCESS, err
}

// 判断是否存在数组对象中
func isExit(arr []ledger.BillChildren, time string) (bool, int) {
	for i, v := range arr {
		if v.Time == time {
			return true, i
		}
	}
	return false, 9999
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
