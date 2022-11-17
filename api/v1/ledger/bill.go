package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
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

}

// GetBillList 获取账单列表
// @Summary 获取账单列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.Bill}
// @Router /ledger/bill/list [get]
func (*ApiLedger) GetBillList(c *gin.Context) {

}
