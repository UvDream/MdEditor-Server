package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
)

// CreateLedgerCategory 创建账本分类
// @Summary 创建账本分类
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.LedgerCategory true "创建账单分类"
// @Success 200 {object} code.Response{data=ledger.LedgerCategory,code=int,msg=string,success=bool}
// @Router /ledger/category/create [post]
func (*ApiLedger) CreateLedgerCategory(c *gin.Context) {

}

// DeleteLedgerCategory 删除账本分类
// @Summary 删除账本分类
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query int true "账本分类ID"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool}
// @Router /ledger/category/delete [delete]
func (*ApiLedger) DeleteLedgerCategory(c *gin.Context) {

}

// UpdateLedgerCategory 更新账本分类
// @Summary 更新账本分类
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.LedgerCategory true "更新账单分类"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.LedgerCategory}
// @Router /ledger/category/update [put]
func (*ApiLedger) UpdateLedgerCategory(c *gin.Context) {

}

// GetLedgerCategoryList 获取账本分类列表
// @Summary 获取账本分类列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query int true "账本分类ID"
// @Param types query string true "分类类型"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.LedgerCategory}
// @Router /ledger/category/list [get]
func (*ApiLedger) GetLedgerCategoryList(c *gin.Context) {
	id := c.Query("id")
	types := c.Query("types")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := ledgerService.GetLedgerCategoryList(id, types)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
