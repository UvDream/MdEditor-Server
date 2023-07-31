package ledger

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	keyWord := c.Query("key_word")
	//TODO 查询是否有权限获取
	//先查询账本是否存在
	if err := db.Where("id = ?", query.LedgerID).First(&ledger.Ledger{}).Error; err != nil {
		return data, total, code.ErrorGetLedger, err
	}
	//时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	//查询备注
	if query.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+query.Remark+"%")
	}
	//查询金额
	if query.Amount != "" {
		db = db.Where("amount = ?", query.Amount)
	}
	if keyWord != "" {
		//	查询备注或者金额
		db = db.Where("remark LIKE ? OR amount = ?", "%"+keyWord+"%", keyWord)
	}
	//查询总数
	if err := db.Model(&ledger.Bill{}).Where("ledger_id = ?", query.LedgerID).Count(&total).Error; err != nil {
		return data, total, code.ErrorGetBill, err
	}
	if query.Sort == "" {
		//	如果为空就根据创建时间排序
		query.Sort = "create_time"
	}
	//查询
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Where("ledger_id = ?", query.LedgerID).Order(query.Sort + " desc").Find(&bill).Error; err != nil {
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

// GetBillNormalListService 获取常用账单列表
func (*LedgersService) GetBillNormalListService(query ledger.BillRequest, userID string, c *gin.Context) (data []ledger.Bill, total int64, cd int, income float64, expenditure float64, err error) {
	db := global.DB
	dp := global.DB
	di := global.DB
	de := global.DB
	if err := dp.Where("id = ?", query.LedgerID).First(&ledger.Ledger{}).Error; err != nil {
		return data, total, code.ErrorGetLedger, income, expenditure, err
	}
	//时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
		di = di.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
		de = de.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	keyWord := c.Query("key_word")
	if keyWord != "" {
		db = db.Where("remark LIKE ? OR amount = ?", "%"+keyWord+"%", keyWord)
		di = di.Where("remark LIKE ? OR amount = ?", "%"+keyWord+"%", keyWord)
		de = de.Where("remark LIKE ? OR amount = ?", "%"+keyWord+"%", keyWord)
	}
	//是否计入收支
	if query.NotBudget != "" {
		db = db.Where("budget_id = ?", query.NotBudget)
	}
	//支出/收入
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}
	ledgerTotal := &struct {
		IncomeTotal float64 `json:"income_total"`
		ExpendTotal float64 `json:"expend_total"`
	}{}

	//排序
	if query.Sort == "" {
		query.Sort = "create_time"
	}
	//查询
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Where("ledger_id = ?", query.LedgerID).Order(query.Sort + " desc").Find(&data).Error; err != nil {
		return data, total, code.ErrorGetBill, income, expenditure, err
	}
	for i, v := range data {
		data[i].CategoryName, err = getCategoryName(v.CategoryID)
		if err != nil {
			return data, total, code.ErrorGetCategory, income, expenditure, err
		}
	}

	//查询总收入
	if err := di.Model(&ledger.Bill{}).Where("type = ?", "1").Select("sum(amount) as income_total").Scan(&ledgerTotal).Error; err != nil {
		return data, total, code.ErrorGetBill, ledgerTotal.IncomeTotal, ledgerTotal.ExpendTotal, err
	}
	//查询总支出
	if err := de.Model(&ledger.Bill{}).Where("type = ?", "0").Select("sum(amount) as expend_total").Scan(&ledgerTotal).Error; err != nil {
		return data, total, code.ErrorGetBill, ledgerTotal.IncomeTotal, ledgerTotal.ExpendTotal, err
	}
	//查询总数
	if err := db.Model(&ledger.Bill{}).Where("ledger_id = ?", query.LedgerID).Count(&total).Error; err != nil {
		return data, total, code.ErrorGetBill, income, expenditure, err
	}
	return data, total, code.SUCCESS, ledgerTotal.IncomeTotal, ledgerTotal.ExpendTotal, err
}

// 根据分类ID获取分类下的名称
func getCategoryName(id string) (name string, err error) {
	db := global.DB
	var category ledger.CategoryLedger
	//首先差是否有parentID
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		return name, err
	}
	if category.ParentID == "" {
		return category.Name, nil
	}
	//	根据id查询父级
	var parent ledger.CategoryLedger
	if err := db.Where("id = ?", category.ParentID).First(&parent).Error; err != nil {
		return name, err
	}
	//	返回父级名称+子级名称
	return parent.Name + "-" + category.Name, nil
}

func (*LedgersService) GetBillListByCategoryIdService(query ledger.BillRequest, userId string, c *gin.Context) (
	data []ledger.Bill, cd int, err error) {
	db := global.DB
	dc := global.DB
	//排序
	if query.Sort != "" {
		db = db.Order(query.Sort + " desc")
	} else {
		db = db.Order("create_time desc")
	}

	// 账单类型
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}
	//是否计入收支
	if query.NotBudget != "" {
		db = db.Where("budget_id = ?", query.NotBudget)
	}
	//时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	//查询列表
	if err := db.Preload(clause.Associations).Find(&data).Error; err != nil {
		return data, code.ErrorGetBill, err
	}
	for i, v := range data {
		data[i].CategoryName, err = getCategoryName(v.CategoryID)
		if err != nil {
			return data, code.ErrorGetCategory, err
		}
	}
	if query.CategoryID == "" {
		return data, code.SUCCESS, err
	}
	//查询分类
	var category ledger.CategoryLedger
	//根据Id查询分类并且查询子级
	if err := dc.Where("id = ?", query.CategoryID).Preload(clause.Associations).First(&category).Error; err != nil {
		return data, code.ErrorGetCategory, err
	}

	//如果没有父级就查询子级
	if category.ParentID == "" {
		FindChildren(dc, category.ID, &category.Children)
	}
	var newData []ledger.Bill
	//如果没有父级直接筛选
	if category.ParentID == "" {
		for i, v := range data {
			if v.CategoryID == category.ID {
				newData = append(newData, data[i])
			}
		}
	}
	//如果有父级就筛选子级
	if category.ParentID != "" {
		for i, v := range data {
			if v.CategoryID == category.ID {
				newData = append(newData, data[i])
			}
			for _, j := range category.Children {
				if v.CategoryID == j.ID {
					//查询分类名称
					data[i].CategoryName, err = getCategoryName(v.CategoryID)
					if err != nil {
						return data, code.ErrorGetCategory, err
					}
					newData = append(newData, data[i])
				}
			}
		}
	}
	return newData, code.SUCCESS, err
}

func FindChildren(db *gorm.DB, parentID string, category *[]ledger.CategoryLedger) error {
	if err := db.Where("parent_id = ?", parentID).Find(&category).Error; err != nil {
		return err
	}
	return nil
}
