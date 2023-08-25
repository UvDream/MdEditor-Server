package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
)

// GetIconList 获取图标列表
// @Summary 获取图标列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param ledger_id query int true "账单ID"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool}
// @Router /ledger/icon/list [get]
func (*ApiLedger) GetIconList(c *gin.Context) {
	ledgerId := c.Query("ledger_id")
	data, cd, err := ledgerService.GetIconListService(ledgerId)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// GetIconColorList 获取图标颜色列表
// @Summary 获取图标颜色列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param ledger_id query int true "账单ID"
// @Param is_bg_color query string false "是否是背景颜色"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool}
// @Router /ledger/icon/color/list [get]
func (*ApiLedger) GetIconColorList(c *gin.Context) {
	ledgerId := c.Query("ledger_id")
	isBgColor := c.Query("is_bg_color")
	data, cd, err := ledgerService.GetIconColorListService(ledgerId, isBgColor)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
