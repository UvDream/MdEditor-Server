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

// GetMemberStatisticsService  获取成员统计
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

// GetCategoryStatisticsDetailService 获取分类详细统计
func (*LedgersService) GetCategoryStatisticsDetailService(filter ledger.CategoryDetailStatisticsData) (data []ledger.CategoryStatisticsData, cd int, err error) {
	db := global.DB
	var categoryList []ledger.CategoryLedger
	if filter.CategoryID == "" {
		//	查询大类
		if err := db.Where("ledger_id = ? and parent_id = ?", filter.LedgerID, "").Find(&categoryList).Error; err != nil {
			return nil, code.ErrorGetCategoryStatisticsDetail, err
		}
		for _, v := range categoryList {
			var itemData ledger.CategoryStatisticsData
			itemData.CategoryID = v.ID
			itemData.CategoryName = v.Name
			//查询大类账单总额
			amount, cd, err := getCategoryAmount(v.ID, filter.LedgerID, filter.StartTime, filter.EndTime, filter.Type)
			if err != nil {
				return nil, cd, err
			}
			itemData.Amount = amount
			//	查询小类
			arr, cd, err := getCategoryList(v.ID, filter.LedgerID)
			if err != nil {
				return nil, cd, err
			}

			for _, b := range arr {
				//	查询小类账单总额
				amount, cd, err := getCategoryAmount(b.ID, filter.LedgerID, filter.StartTime, filter.EndTime, filter.Type)
				if err != nil {
					return nil, cd, err
				}
				itemData.Amount += amount
			}
			data = append(data, itemData)
		}
	} else {
		//	存在大类ID
		//	查询所有的小类
		arr, cd, err := getCategoryList(filter.CategoryID, filter.LedgerID)
		if err != nil {
			return nil, cd, err
		}
		for _, b := range arr {
			//	查询小类账单总额
			amount, cd, err := getCategoryAmount(b.ID, filter.LedgerID, filter.StartTime, filter.EndTime, filter.Type)
			if err != nil {
				return nil, cd, err
			}
			var itemData ledger.CategoryStatisticsData
			itemData.CategoryID = b.ID
			itemData.CategoryName = b.Name
			itemData.Amount = amount
			data = append(data, itemData)
		}
	}

	return data, code.SUCCESS, err
}

// 根据大类ID查询小类数据
func getCategoryList(categoryID string, ledgerID string) (data []ledger.CategoryLedger, cd int, err error) {
	db := global.DB
	if err := db.Where("ledger_id = ? and parent_id = ?", ledgerID, categoryID).Find(&data).Error; err != nil {
		return nil, code.ErrorGetCategoryStatisticsDetail, err
	}
	return data, code.SUCCESS, err
}

// 根据小类查询该小类账单总额
func getCategoryAmount(categoryID string, ledgerID string, startTime string, endTime string, types string) (data float64, cd int, err error) {
	db := global.DB
	db = db.Model(&ledger.Bill{})
	if startTime != "" {
		db = db.Where("create_time >?", startTime)
	}
	if endTime != "" {
		db = db.Where("create_time < ?", endTime)
	}

	if err := db.Where("ledger_id = ? and category_id = ? and type = ?", ledgerID, categoryID, types).Select("COALESCE(sum(amount),0) as amount").Scan(&data).Error; err != nil {
		return 0, code.ErrorGetCategoryStatisticsDetail, err
	}
	return data, code.SUCCESS, err
}

// GetTotalAmountService 获取总金额
func (*LedgersService) GetTotalAmountService(filter ledger.CategoryDetailStatisticsData) (data float64, cd int, err error) {
	db := global.DB
	db = db.Model(&ledger.Bill{})
	if filter.StartTime != "" {
		db = db.Where("create_time >?", filter.StartTime)
	}
	if filter.EndTime != "" {
		db = db.Where("create_time < ?", filter.EndTime)
	}
	if err := db.Where("ledger_id = ? and type = ?", filter.LedgerID, filter.Type).Select("COALESCE(sum(amount),0) as amount").Scan(&data).Error; err != nil {
		return 0, code.ErrorGetTotalAmount, err
	}
	return data, code.SUCCESS, err
}
