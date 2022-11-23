package ledger

import (
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/ledger"
	"strconv"
)

func (*LedgersService) CreateBudget(budget ledger.MoneyBudget) (data ledger.MoneyBudget, cd int, err error) {
	db := global.DB
	if budget.IsSameMonth {
		var budgets []ledger.MoneyBudget
		//十二个月
		for i := 0; i < 12; i++ {
			budget.Date = i + 1
			budgets = append(budgets, budget)
		}
		if err = db.Create(&budgets).Error; err != nil {
			return data, code.ErrCreateBudget, err
		}
	} else {
		if err := db.Create(&budget).Error; err != nil {
			return data, code.ErrCreateBudget, err
		}
	}
	return data, code.SUCCESS, nil
}
func (*LedgersService) DeleteBudget(id string) (cd int, err error) {
	db := global.DB
	if err := db.Delete(&ledger.MoneyBudget{}, id).Error; err != nil {
		return code.ErrDeleteBudget, err
	}
	return code.SUCCESS, nil
}
func (*LedgersService) UpdateBudget(budget ledger.MoneyBudget) (data ledger.MoneyBudget, cd int, err error) {
	db := global.DB
	//查询是否存在
	if err := db.First(&data, budget.ID).Error; err != nil {
		return data, code.ErrBudgetNotExist, err
	}
	if budget.IsSameMonth {
		var budgets []ledger.MoneyBudget
		//十二个月
		for i := 0; i < 12; i++ {
			budget.Date = i + 1
			budgets = append(budgets, budget)
		}
		if err = db.Save(&budgets).Error; err != nil {
			return data, code.ErrCreateBudget, err
		}
	} else {
		if err := db.Save(&budget).Error; err != nil {
			return data, code.ErrUpdateBudget, err
		}
	}
	return data, code.SUCCESS, nil
}

func (*LedgersService) GetBudgetList(ledgerID string, year string) (data []ledger.MoneyBudget, cd int, err error) {
	if err := global.DB.Where("ledger_id = ?", ledgerID).Where("year = ?", year).Preload(clause.Associations).Order("date").Find(&data).Error; err != nil {
		return data, code.ErrGetBudget, err
	}
	//算出每个月支出
	for i, k := range data {
		//	时间
		year := strconv.Itoa(k.Year)
		month := strconv.Itoa(k.Date)
		if k.Date < 10 {
			month = "0" + month
		}
		currentMonth := year + "-" + month + "-01"
		nextMonth := year + "-" + strconv.Itoa(k.Date+1) + "-01"
		//	查询
		d := &struct {
			Expenditure float64 `json:"expenditure"`
		}{}
		if err := global.DB.Model(&ledger.Bill{}).Where("type = ?", 0).Where("ledger_id = ?", ledgerID).Where("create_time BETWEEN ? AND ?", currentMonth, nextMonth).Select("sum(amount) as expenditure").Scan(&d).Error; err != nil {
			return data, code.ErrGetBudget, err
		}

		data[i].Expenditure = d.Expenditure
	}
	return data, code.SUCCESS, nil
}
