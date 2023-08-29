package admin

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
)

type LedgerAdminApi struct {
}

// GetLedgerList 获取账本列表admin
// @Summary 获取账本列表admin
// @Description 获取账本列表admin
// @Tags admin/ledger
// @Produce  json
// @Param query query models.PaginationRequest true "分页参数"
// @Param query query ledger.LedgerRequest true "搜索参数"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.Ledger}
// @Router /admin/ledger/list [get]
func (*LedgerAdminApi) GetLedgerList(c *gin.Context) {
	var query ledger.LedgerRequest
	err := c.ShouldBindQuery(&query)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	data, total, cd, err := ledgerAdminService.GetLedgerListService(query, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}
