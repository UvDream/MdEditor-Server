package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
)

// GetCalendar 账单控制器
// @Summary 账单控制器
// @Tags statistics
// @Accept  json
// @Produce  json
// @Param year query string true "年份"
// @Param month query string true "月份"
// @Param ledger_id query string true "账本id"
// @Success 200 {object} code.Response{data=ledger.CalendarData,code=int,msg=string,success=bool}
// @Router /ledger/statistics/calendar [get]
func (*ApiLedger) GetCalendar(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")
	ledgerID := c.Query("ledger_id")
	data, cd, err := ledgerService.GetCalendarService(year, month, ledgerID)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// GetHomeStatistics 首页统计
// @Summary 首页统计
// @Tags statistics
// @Accept  json
// @Produce  json
// @Param ledger_id query string true "账本id"
// @Param start_time query string true "开始时间"
// @Param end_time query string true "结束时间"
// @Success 200 {object} code.Response{data=ledger.HomeStatisticsData,code=int,msg=string,success=bool}
// @Router /ledger/statistics/home [get]
func (*ApiLedger) GetHomeStatistics(c *gin.Context) {
	ledgerID := c.Query("ledger_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	data, cd, err := ledgerService.GetHomeStatisticsService(ledgerID, startTime, endTime)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
