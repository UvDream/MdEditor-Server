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
