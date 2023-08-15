package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/ledger"
	"server/utils"
)

func (*LedgerAdminService) GetBillListService(query ledger.BillRequest, c *gin.Context) (data []ledger.Bill, total int64, cd int, income float64, expenditure float64, err error) {
	db := global.DB
	di := global.DB
	de := global.DB
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
	//amount
	if query.Amount != "" {
		db = db.Where("amount = ?", query.Amount)
		di = di.Where("amount = ?", query.Amount)
		de = de.Where("amount = ?", query.Amount)
	}
	//remark
	if query.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+query.Remark+"%")
		di = di.Where("remark LIKE ?", "%"+query.Remark+"%")
		de = de.Where("remark LIKE ?", "%"+query.Remark+"%")
	}
	//是否计入收支
	if query.NotBudget != "" {
		db = db.Where("not_budget = ?", query.NotBudget)
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
	//查询多分类
	if len(query.CategoryIDs) > 0 {
		db = db.Where("category_id in (?)", query.CategoryIDs)
		di = di.Where("category_id in (?)", query.CategoryIDs)
		de = de.Where("category_id in (?)", query.CategoryIDs)
	}
	//查询
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Order(query.Sort + " desc").Find(&data).Error; err != nil {
		return data, total, code.ErrorGetBill, income, expenditure, err
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
	if err := db.Model(&ledger.Bill{}).Count(&total).Error; err != nil {
		return data, total, code.ErrorGetBill, income, expenditure, err
	}
	return data, total, code.SUCCESS, income, expenditure, nil
}
