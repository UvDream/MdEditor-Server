package ledger

import (
	"server/code"
	"server/global"
	"server/models/ledger"
	"server/utils"
)

func (*LedgersService) GetCategoryStatisticsService(ledgerID string, startTime string, endTime string, types string) (data []ledger.CategoryStatisticsData, cd int, err error) {
	var category []ledger.CategoryLedger
	db := global.DB
	//查出账本所有分类
	if err := db.Where("ledger_id = ?", ledgerID).Find(&category).Error; err != nil {
		return nil, code.ErrorGetCategory, err
	}
	//求出账本支出/收入总额
	ledgerTotal := &struct {
		IncomeTotal float64 `json:"income_total"`
		ExpendTotal float64 `json:"expend_total"`
	}{}
	if types == "0" {
		if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? and create_time between ? and ?", ledgerID, types, startTime, endTime).Select("sum(amount) as expend_total").Scan(ledgerTotal).Error; err != nil {
			return nil, code.ErrorGetCategory, err
		}
	} else {
		if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? and create_time between ? and ?", ledgerID, types, startTime, endTime).Select("sum(amount) as income_total").Scan(ledgerTotal).Error; err != nil {
			return nil, code.ErrorGetCategory, err
		}
	}
	//查出分类的总数量
	for _, v := range category {
		d := &struct {
			Income      float64 `json:"income"`
			Expenditure float64 `json:"expenditure"`
		}{}
		if types == "0" {
			if err := db.Model(&ledger.Bill{}).Where("category_id = ? and type = ? and create_time BETWEEN ? AND ?", v.ID, types, startTime, endTime).Select("sum(amount) as expenditure").Scan(d).Error; err != nil {
				return nil, code.ErrorGetCategoryStatistics, err
			}
		} else {
			if err := db.Model(&ledger.Bill{}).Where("category_id = ? and type = ? and create_time BETWEEN ? AND ?", v.ID, types, startTime, endTime).Select("sum(amount) as income").Scan(d).Error; err != nil {
				return nil, code.ErrorGetCategoryStatistics, err
			}
		}
		var categoryStatisticsData ledger.CategoryStatisticsData
		categoryStatisticsData.CategoryName = v.Name
		categoryStatisticsData.CategoryID = v.ID
		if types == "0" {
			categoryStatisticsData.Amount = d.Expenditure
			if d.Expenditure == 0 || ledgerTotal.ExpendTotal == 0 {
				categoryStatisticsData.Ratio = 0
			} else {
				categoryStatisticsData.Ratio = d.Expenditure / ledgerTotal.ExpendTotal
			}
		} else {
			categoryStatisticsData.Amount = d.Income
			if d.Income == 0 || ledgerTotal.IncomeTotal == 0 {
				categoryStatisticsData.Ratio = 0
			} else {
				categoryStatisticsData.Ratio = d.Income / ledgerTotal.IncomeTotal
			}
		}
		data = append(data, categoryStatisticsData)
	}
	return data, code.SUCCESS, err
}

func (*LedgersService) GetIncomeExpenditureStatisticsService(ledgerID string, startTime string, endTime string, types string, isYear string) (data []ledger.IncomeExpenditureStatisticsData, cd int, err error) {
	db := global.DB
	dayList, _ := utils.GetDateList(startTime, endTime, isYear)
	//求出每天的收入/支出
	for _, v := range dayList {
		d := &struct {
			Income      float64 `json:"income"`
			Expenditure float64 `json:"expenditure"`
		}{}
		if types == "0" {
			if isYear == "0" {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 0, 1)).Select("sum(amount) as expenditure").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			} else {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 1, 0)).Select("sum(amount) as expenditure").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			}
		} else {
			if isYear == "0" {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 0, 1)).Select("sum(amount) as income").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			} else {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 1, 0)).Select("sum(amount) as income").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			}

		}
		var incomeExpenditureStatisticsData ledger.IncomeExpenditureStatisticsData
		incomeExpenditureStatisticsData.Date = v.Format("2006-01-02")
		if types == "0" {
			incomeExpenditureStatisticsData.Amount = d.Expenditure
		} else {
			incomeExpenditureStatisticsData.Amount = d.Income
		}
		data = append(data, incomeExpenditureStatisticsData)
	}
	return data, code.SUCCESS, err
}
