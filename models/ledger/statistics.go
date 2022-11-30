package ledger

type CategoryStatisticsData struct {
	//	分类名称
	CategoryName string `json:"category_name"`
	//	分类ID
	CategoryID string `json:"category_id"`
	//	分类金额
	Amount float64 `json:"amount"`
	//	账本总支出/收入占比
	Ratio float64 `json:"ratio"`
}

type IncomeExpenditureStatisticsData struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}
type MemberStatisticsData struct {
	//	成员名称
	Name string `json:"name"`
	//	成员ID
	MemberID string `json:"member_id"`
	//	成员金额
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
}

type PersonalStatisticsData struct {
	//	打卡天数
	AttendanceDays int64 `json:"attendance_days"`
	//	记账天数
	AccountingDays int64 `json:"accounting_days"`
	//	记账数目
	AccountingNumber int64 `json:"accounting_number"`
	//记账总额
	AccountingTotal float64 `json:"accounting_total"`
}
