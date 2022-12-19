package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
	"server/utils"
	"time"
)

// CreateBill 创建账单
// @Summary 创建账单
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.Bill true "创建账单"
// @Success 200 {object} code.Response{data=ledger.Bill,code=int,msg=string,success=bool}
// @Router /ledger/bill/create [post]
func (*ApiLedger) CreateBill(c *gin.Context) {
	var bill ledger.Bill
	err := c.ShouldBindJSON(&bill)

	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	local, _ := time.LoadLocation("Local")
	bill.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", bill.Date, local)
	if bill.LedgerID == "" {
		code.FailResponse(code.ErrorMissingLedgerId, c)
		return
	}
	userID := utils.FindUserID(c)
	bill.CreatorID = userID
	data, cd, err := ledgerService.AddBillService(bill)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	code.SuccessResponse(data, cd, c)

}

// DeleteBill 删除账单
// @Summary 删除账单
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query int true "账单ID"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool}
// @Router /ledger/bill/delete [delete]
func (*ApiLedger) DeleteBill(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerService.DeleteBillService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, code.SUCCESS, c)
}

// UpdateBill 更新账单
// @Summary 更新账单
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.Bill true "更新账单"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.Bill}
// @Router /ledger/bill/update [put]
func (*ApiLedger) UpdateBill(c *gin.Context) {
	var bill ledger.Bill
	err := c.ShouldBindJSON(&bill)

	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	local, _ := time.LoadLocation("Local")
	bill.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", bill.Date, local)
	if bill.LedgerID == "" {
		code.FailResponse(code.ErrorMissingLedgerId, c)
		return
	}
	if bill.ID == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := ledgerService.UpdateBillService(bill)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)

}

// GetBillList 获取账单列表
// @Summary 获取账单列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param  query  query    ledger.BillRequest  true  "参数"
// @Param  query  query    models.PaginationRequest  true  "参数"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.BillChildren}
// @Router /ledger/bill/list [get]
func (*ApiLedger) GetBillList(c *gin.Context) {
	var query ledger.BillRequest
	err := c.ShouldBindQuery(&query)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.FindUserID(c)
	data, total, cd, err := ledgerService.GetBillListService(query, userID, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(gin.H{
		"list":  data,
		"total": total,
	}, cd, c)
}

// GetBillDetail 获取账单详情
// @Summary 获取账单详情
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query int true "账单ID"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.Bill}
// @Router /ledger/bill/detail [get]
func (*ApiLedger) GetBillDetail(c *gin.Context) {
	id := c.Query("id")
	data, cd, err := ledgerService.GetBillDetailService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
