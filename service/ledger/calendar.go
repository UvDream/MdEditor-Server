package ledger

import (
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
	var bills []ledger.Bill
	if err := db.Model(&ledger.Bill{}).Where("ledger_id = ?", ledgerID).Where("create_time BETWEEN ? AND ?", year+"-"+month+"-01", year+"-"+month+"-"+strconv.Itoa(day)).Find(&bills).Error; err != nil {
		return nil, code.ErrorGetBill, err
	}
	// 按照每天生成日历数据
	firstDay, _ := time.Parse("2006-01-02", year+"-"+month+"-01")
	for i := 1; i <= day; i++ {
		data = append(data, CalendarData{
			Day: firstDay.AddDate(0, 0, i-1).Format("2006-01-02"),
		})
	}
	// 遍历账单，将账单数据填充到日历数据中
	for _, bill := range bills {
		for i, v := range data {
			if v.Day == bill.CreateTime.Format("2006-01-02") {
				if bill.Type == "1" {
					data[i].Income += bill.Amount
				} else {
					data[i].Expenditure += bill.Amount
				}
			}
		}
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
