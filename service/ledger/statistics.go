package ledger

import (
	"server/code"
	"server/global"
	"server/models/ledger"
	"server/utils"
	"time"
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
		//查询是否有ParentID不等于空
		var parentCategory ledger.CategoryLedger

		if v.ParentID != "" {
			//	查询父级分类
			if err := db.Where("id = ?", v.ParentID).Find(&parentCategory).Error; err != nil {
				return nil, code.ErrorGetCategoryStatistics, err
			}
			categoryStatisticsData.CategoryName = parentCategory.Name + "-" + v.Name
		} else {
			categoryStatisticsData.CategoryName = v.Name
		}

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
		var d ledger.DataSummary
		nextDay := v.AddDate(0, 0, 1)
		nextMonth := v.AddDate(0, 1, 0)
		if types == "0" {
			if isYear == "0" {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >= ?", v).Where("create_time < ?", nextDay).Select("sum(amount) as expenditure").Scan(&d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			} else {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >= ?", v).Where("create_time < ?", nextMonth).Select("sum(amount) as expenditure").Scan(&d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			}
		} else {
			if isYear == "0" {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >= ?", v).Where("create_time < ?", nextDay).Select("sum(amount) as income").Scan(&d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			} else {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? ", ledgerID, types).Where("create_time >= ?", v).Where("create_time < ?", nextMonth).Select("sum(amount) as income").Scan(&d).Error; err != nil {
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

func (*LedgersService) GetMemberStatisticsService(ledgerID string, startTime string, endTime string, types string, isYear string) (data []ledger.MemberStatisticsData, cd int, err error) {
	db := global.DB
	dayList, _ := utils.GetDateList(startTime, endTime, isYear)
	//先查出用户
	var ledgerList ledger.Ledger
	if err := db.Where("id = ?", ledgerID).Preload("User").Preload("Creator").First(&ledgerList).Error; err != nil {
		return nil, code.ErrorGetMemberStatistics, err
	}
	ledgerList.User = append(ledgerList.User, ledgerList.Creator)
	for _, v := range ledgerList.User {
		arr, cd, err := getAmount(ledgerID, dayList, types, isYear, v.ID, v.NickName)
		if err != nil {
			return nil, cd, err
		}
		data = append(data, arr...)
	}
	return data, code.SUCCESS, err
}

type UserStatisticsData struct {
	Income      float64 `json:"income"`
	Expenditure float64 `json:"expenditure"`
}

func getAmount(ledgerID string, dateList []time.Time, types string, isYear string, userID string, nickName string) (data []ledger.MemberStatisticsData, cd int, err error) {
	db := global.DB
	for _, v := range dateList {
		d := &struct {
			Income      float64 `json:"income"`
			Expenditure float64 `json:"expenditure"`
		}{}
		if types == "0" {
			if isYear == "0" {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? and creator_id = ?", ledgerID, types, userID).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 0, 1)).Select("sum(amount) as expenditure").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			} else {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? and creator_id = ?", ledgerID, types, userID).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 1, 0)).Select("sum(amount) as expenditure").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			}
		} else {
			if isYear == "0" {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? and creator_id = ?", ledgerID, types, userID).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 0, 1)).Select("sum(amount) as income").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			} else {
				if err := db.Model(&ledger.Bill{}).Where("ledger_id = ? and type = ? and creator_id = ?", ledgerID, types, userID).Where("create_time >?", v).Where("create_time < ?", v.AddDate(0, 1, 0)).Select("sum(amount) as income").Scan(d).Error; err != nil {
					return nil, code.ErrorGetIncomeExpenditureStatistics, err
				}
			}

		}
		var memberStatisticsData ledger.MemberStatisticsData
		memberStatisticsData.Date = v.Format("2006-01-02")
		memberStatisticsData.Name = nickName
		memberStatisticsData.MemberID = userID
		if types == "0" {
			memberStatisticsData.Amount = d.Expenditure
		} else {
			memberStatisticsData.Amount = d.Income
		}
		data = append(data, memberStatisticsData)
	}
	return data, code.SUCCESS, err
}
