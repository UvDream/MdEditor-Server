package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
	"server/utils"
)

// CreateLedgerCategory 创建账本分类
// @Summary 创建账本分类
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.CategoryLedger true "创建账单分类"
// @Success 200 {object} code.Response{data=ledger.LedgerCategory,code=int,msg=string,success=bool}
// @Router /ledger/category/create [post]
func (*ApiLedger) CreateLedgerCategory(c *gin.Context) {
	var category ledger.CategoryLedger
	if err := c.ShouldBindJSON(&category); err != nil {
		code.FailResponse(code.ErrorUpdatePasswordMissingParam, c)
		return
	}
	userID := utils.FindUserID(c)
	category.UserID = userID
	data, cd, err := ledgerService.CreateLedgerCategory(category)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
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
	id := c.Query("id")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	cd, err := ledgerService.DeleteLedgerCategory(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// UpdateLedgerCategory 更新账本分类
// @Summary 更新账本分类
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.CategoryLedger true "更新账单分类"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.LedgerCategory}
// @Router /ledger/category/update [put]
func (*ApiLedger) UpdateLedgerCategory(c *gin.Context) {
	var category ledger.CategoryLedger
	if err := c.ShouldBindJSON(&category); err != nil {
		code.FailResponse(code.ErrorUpdatePasswordMissingParam, c)
		return
	}
	//TODO: 验证用户是否有权限
	//userID := utils.FindUserID(c)
	//if userID!=category.UserID{
	//	code.FailResponse(code.ErrorNoPermission, c)
	//	return
	//}
	data, cd, err := ledgerService.UpdateLedgerCategory(category)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)

}

// GetLedgerCategoryList 获取账本分类列表
// @Summary 获取账本分类列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query string true "账本分类ID"
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

//GetLedgerCategoryDetail 获取账本分类详情
// @Summary 获取账本分类详情
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query string true "账本分类ID"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.LedgerCategory}
// @Router /ledger/category/detail [get]
func (*ApiLedger) GetLedgerCategoryDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := ledgerService.GetLedgerCategoryDetail(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
