package ledger

import (
	"fmt"
	"server/code"
	"server/global"
	"server/models/ledger"
	"strconv"
	"time"
)

type CalendarData struct {
	Day         string  `json:"day"`
	Income      float64 `json:"income"`
	Expenditure float64 `json:"expenditure"`
}

func (*LedgersService) GetCalendarService(year string, month string, ledgerID string) (data []CalendarData, cd int, err error) {
	db := global.DB
	intYear, err := strconv.Atoi(year)
	if err != nil {
		return nil, code.ERROR, err
	}
	intMonth, err := strconv.Atoi(month)
	if err != nil {
		return nil, code.ERROR, err
	}
	day := getYearMonthToDay(intYear, intMonth)
	if intMonth < 10 {
		month = "0" + month
	}
	firstDay, _ := time.Parse("2006/01/02 15:04:05", year+"/"+month+"/01 00:00:00")
	for i := 0; i < day; i++ {
		toDay := firstDay.AddDate(0, 0, i)
		fmt.Println(toDay)
		//	查出当天的收入和支出的总和amount
		type total struct {
			Income      float64 `json:"income"`
			Expenditure float64 `json:"expenditure"`
		}
		var d total
		//算出支出
		if err := db.Model(&ledger.Bill{}).Where("type = ?", 0).Where("ledger_id = ?", ledgerID).Where("create_time BETWEEN ? AND ?", toDay, toDay.AddDate(0, 0, 1)).Select("sum(amount) as expenditure").Scan(&d).Error; err != nil {
			return data, code.ErrorGetBill, err
		}
		if err := db.Model(&ledger.Bill{}).Where("type = ?", 1).Where("ledger_id = ?", ledgerID).Where("create_time BETWEEN ? AND ?", toDay, toDay.AddDate(0, 0, 1)).Select("sum(amount) as income").Scan(&d).Error; err != nil {
			return data, code.ErrorGetBill, err
		}
		data = append(data, CalendarData{
			Day:         toDay.Format("2006-01-02"),
			Income:      d.Income,
			Expenditure: d.Expenditure,
		})
	}
	return data, code.SUCCESS, nil
}

func getYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]struct{}{
		1:  struct{}{},
		3:  struct{}{},
		5:  struct{}{},
		7:  struct{}{},
		8:  struct{}{},
		10: struct{}{},
		12: struct{}{},
	}
	if _, ok := day31[month]; ok {
		return 31
	}
	// 有30天的月份
	day30 := map[int]struct{}{
		4:  struct{}{},
		6:  struct{}{},
		9:  struct{}{},
		11: struct{}{},
	}
	if _, ok := day30[month]; ok {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}

func (*LedgersService) GetHomeStatisticsService(ledgerID string, startTime string, endTime string) (data ledger.HomeStatisticsData, cd int, err error) {
	db := global.DB
	//查询当月的收入/支出总和
	d := &struct {
		Income      float64 `json:"income"`
		Expenditure float64 `json:"expenditure"`
	}{}

	if err := db.Model(&ledger.Bill{}).Where("type = ?", 1).Where("ledger_id = ?", ledgerID).Where("create_time BETWEEN ? AND ?", startTime, endTime).Select("sum(amount) as income").Scan(d).Error; err != nil {
		return data, code.ErrorGetBill, err
	}
	if err := db.Model(&ledger.Bill{}).Where("type = ?", 0).Where("ledger_id = ?", ledgerID).Where("create_time BETWEEN ? AND ?", startTime, endTime).Select("sum(amount) as expenditure").Scan(d).Error; err != nil {
		return data, code.ErrorGetBill, err
	}
	data = ledger.HomeStatisticsData{
		Income:      d.Income,
		Expenditure: d.Expenditure,
	}
	return data, code.SUCCESS, nil
}
