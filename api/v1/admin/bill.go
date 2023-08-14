package admin

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
	service "server/service/system"
	"server/utils"
)

// GetBillList 获取账单列表admin
// @Summary 获取账单列表admin
// @Description 获取账单列表admin
// @Tags admin
// @Produce  json
// @Param query query models.PaginationRequest true "分页参数"
// @Param query query ledger.BillRequest true "搜索参数"
// @Success 200 {object} code.Response
// @Router /admin/ledger/bill/list [get]
func (*LedgerAdminApi) GetBillList(c *gin.Context) {
	var query ledger.BillRequest
	err := c.ShouldBindQuery(&query)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.FindUserID(c)
	var userService service.SysUserService
	roles, err := userService.FindUserRoles(userID)
	if err != nil {
		code.FailResponse(code.ErrorGetPermission, c)
		return
	}
	if len(roles) > 0 {
		data, total, cd, income, expenditure, err := ledgerAdminService.GetBillListService(query, c)
		if err != nil {
			code.FailResponse(cd, c)
			return
		}
		code.SuccessResponse(gin.H{
			"list":        data,
			"total":       total,
			"income":      income,
			"expenditure": expenditure,
		}, cd, c)
	} else {
		data, total, cd, income, expenditure, err := ledgerService.GetBillNormalListService(query, userID, c)
		if err != nil {
			code.FailResponse(cd, c)
			return
		}
		code.SuccessResponse(gin.H{
			"list":        data,
			"total":       total,
			"income":      income,
			"expenditure": expenditure,
		}, cd, c)
	}

}
