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
	ledger_id := c.Query("ledger_id")
	data, cd, err := ledgerService.GetCalendarService(year, month, ledger_id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
