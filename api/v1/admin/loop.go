package admin

import (
	"github.com/gin-gonic/gin"
	"server/code"
)

// GetLoopList	获取定时任务列表
// @Summary 获取定时任务列表
// @Tags admin/loop
// @Accept  json
// @Produce  json
// @Param query query models.PaginationRequest true "分页参数"
// @Success 200 {object} code.Response
// @Router /admin/ledger/loop/list [get]
func (*LedgerAdminApi) GetLoopList(c *gin.Context) {
	data, total, cd, err := ledgerAdminService.GetLoopListService(c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}
