package ledger

import (
	"server/code"
	"server/global"
	"server/models/ledger"
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

func (*LedgersService) GetBudgetList(ledgerID string) (data []ledger.MoneyBudget, cd int, err error) {
	if err := global.DB.Where("ledger_id = ?", ledgerID).Find(&data).Error; err != nil {
		return data, code.ErrGetBudget, err
	}
	return data, code.SUCCESS, nil
}
