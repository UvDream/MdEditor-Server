package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/global"
	"server/models/ledger"
	"server/utils"
)

//GetCategoryStatistics 获取分类统计
// @Tags statistics
// @Summary 获取分类统计
// @Description 获取分类统计
// @Accept  json
// @Produce  json
// @Param ledger_id query string true "账本ID"
// @Param start_time query string true "开始时间"
// @Param end_time query string true "结束时间"
// @Param type query string true "类型"
// @Success 200 {object}  code.Response{data=[]ledger.CategoryStatisticsData,code=int,msg=string,success=bool}
// @Router /ledger/statistics/category [get]
func (*ApiLedger) GetCategoryStatistics(c *gin.Context) {
	ledgerID := c.Query("ledger_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	types := c.Query("type")
	data, cd, err := ledgerService.GetCategoryStatisticsService(ledgerID, startTime, endTime, types)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

//GetIncomeExpenditureStatistics 获取收支统计
// @Tags statistics
// @Summary 获取收支统计
// @Description 获取收支统计
// @Accept  json
// @Produce  json
// @Param ledger_id query string true "账本ID"
// @Param start_time query string true "开始时间"
// @Param end_time query string true "结束时间"
// @Param type query string true "类型"
// @Param is_year query string true "是否按年统计"
// @Success 200 {object}  code.Response{data=[]ledger.IncomeExpenditureStatisticsData,code=int,msg=string,success=bool}
// @Router /ledger/statistics/income_expenditure [get]
func (*ApiLedger) GetIncomeExpenditureStatistics(c *gin.Context) {
	ledgerID := c.Query("ledger_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	types := c.Query("type")
	isYear := c.Query("is_year")
	data, cd, err := ledgerService.GetIncomeExpenditureStatisticsService(ledgerID, startTime, endTime, types, isYear)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// GetMemberStatistics 获取成员统计
// @Tags statistics
// @Summary 获取成员统计
// @Description 获取成员统计
// @Accept  json
// @Produce  json
// @Param ledger_id query string true "账本ID"
// @Param start_time query string true "开始时间"
// @Param end_time query string true "结束时间"
// @Param type query string true "类型"
// @Param is_year query string true "是否按年统计"
// @Success 200 {object}  code.Response{data=[]ledger.MemberStatisticsData,code=int,msg=string,success=bool}
// @Router /ledger/statistics/member [get]
func (*ApiLedger) GetMemberStatistics(c *gin.Context) {
	ledgerID := c.Query("ledger_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	types := c.Query("type")
	isYear := c.Query("is_year")
	data, cd, err := ledgerService.GetMemberStatisticsService(ledgerID, startTime, endTime, types, isYear)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

//GetPersonalStatistics 获取个人统计
// @Tags statistics
// @Summary 获取个人统计
// @Description 获取个人统计
// @Accept  json
// @Produce  json
// @Success 200 {object}  code.Response{data=[]ledger.PersonalStatisticsData,code=int,msg=string,success=bool}
// @Router /ledger/statistics/personal [get]
func (*ApiLedger) GetPersonalStatistics(c *gin.Context) {
	personalStatisticsData := &ledger.PersonalStatisticsData{}
	db := global.DB
	userID := utils.FindUserID(c)
	//记账数目
	if err := db.Model(&ledger.Bill{}).Where("creator_id= ? ", userID).Count(&personalStatisticsData.AccountingNumber).Error; err != nil {
		code.FailResponse(code.ErrorGetPersonalStatistics, c)
		return
	}
	//	记账天数
	if err := db.Model(&ledger.Bill{}).Where("creator_id= ? ", userID).Distinct("create_time").Count(&personalStatisticsData.AccountingDays).Error; err != nil {
		code.FailResponse(code.ErrorGetPersonalStatistics, c)
		return
	}

	//	记账总金额
	if err := db.Model(&ledger.Bill{}).Where("creator_id= ? ", userID).Select("sum(amount) as income").Scan(&personalStatisticsData.AccountingTotal).Error; err != nil {
		code.FailResponse(code.ErrorGetPersonalStatistics, c)
		return
	}
	code.SuccessResponse(personalStatisticsData, code.SUCCESS, c)
}
