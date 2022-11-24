package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
	"server/utils"
)

// CreateBudget 创建预算
// @Tags budget
// @Summary 创建预算
// @accept application/json
// @Produce application/json
// @Param budget body ledger.MoneyBudget true "预算信息"
// @Success 200 {object} code.Response{data=ledger.MoneyBudget}
// @Router  /ledger/budget/create [post]
func (*ApiLedger) CreateBudget(c *gin.Context) {
	var budget ledger.MoneyBudget
	err := c.ShouldBindJSON(&budget)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.FindUserID(c)
	budget.CreatorID = userID
	data, cd, err := ledgerService.CreateBudget(budget)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// DeleteBudget 删除预算
// @Tags budget
// @Summary 删除预算
// @accept application/json
// @Produce application/json
// @Param id query int true "预算ID"
// @Success 200 {object} code.Response{data=ledger.MoneyBudget}
// @Router  /ledger/budget/delete [delete]
func (*ApiLedger) DeleteBudget(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	cd, err := ledgerService.DeleteBudget(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// UpdateBudget 更新预算
// @Tags budget
// @Summary 更新预算
// @accept application/json
// @Produce application/json
// @Param budget body ledger.MoneyBudget true "预算信息"
// @Success 200 {object} code.Response{data=ledger.MoneyBudget}
// @Router  /ledger/budget/update [put]
func (*ApiLedger) UpdateBudget(c *gin.Context) {
	var budget ledger.MoneyBudget
	err := c.ShouldBindJSON(&budget)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	if budget.ID == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := ledgerService.UpdateBudget(budget)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

type BudgetData struct {
	YearBudget  ledger.MoneyBudget   `json:"year_budget"`
	MonthBudget []ledger.MoneyBudget `json:"month_budget"`
}

// GetBudgetList 获取预算
// @Tags budget
// @Summary 获取预算
// @accept application/json
// @Produce application/json
// @Param id query int true "账本ID"
// @Param year query int true "年份"
// @Success 200 {object} code.Response{data=BudgetData}
// @Router  /ledger/budget/list [get]
func (*ApiLedger) GetBudgetList(c *gin.Context) {
	id := c.Query("id")
	year := c.Query("year")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := ledgerService.GetBudgetList(id, year)

	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	yearData, cd, err := ledgerService.GetYearBudget(id, year)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}

	code.SuccessResponse(BudgetData{
		YearBudget:  yearData,
		MonthBudget: data,
	}, cd, c)
}
