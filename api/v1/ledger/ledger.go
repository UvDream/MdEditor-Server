package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/global"
	ledger2 "server/models/ledger"
	"server/utils"
)

type ApiLedger struct{}

// CreateLedger 创建账本
//@Summary 创建账本
//@Tags ledger
//@Accept  json
//@Produce  json
//@Param article body ledger.Ledger true "创建账单"
//@Success 200 {object} code.Response{data=ledger.Ledger,code=int,msg=string,success=bool}
//@Router /ledger/create [post]
func (*ApiLedger) CreateLedger(c *gin.Context) {
	var ledger ledger2.Ledger
	err := c.ShouldBindJSON(&ledger)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	ledger.CreatorID = utils.FindUserID(c)
	data, cd, err := ledgerService.CreateLedger(ledger)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// DeleteLedger 删除账本
//@Summary 删除账本
//@Tags ledger
//@Accept  json
//@Produce  json
//@Param id query int true "账本ID"
//@Success 200 {object} code.Response{code=int,msg=string,success=bool}
//@Router /ledger/delete [delete]
func (*ApiLedger) DeleteLedger(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	cd, err := ledgerService.DeleteLedger(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// UpdateLedger 更新账本
//@Summary 更新账本
//@Tags ledger
//@Accept  json
//@Produce  json
//@Param article body ledger.Ledger true "更新账单"
//@Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.Ledger}
//@Router /ledger/update [put]
func (*ApiLedger) UpdateLedger(c *gin.Context) {
	var ledger ledger2.Ledger
	if err := c.ShouldBindJSON(&ledger); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	db := global.DB
	//	查询账本是否存在
	if err := db.Where("id = ?", ledger.ID).First(&ledger2.Ledger{}).Error; err != nil {
		code.FailResponse(code.ErrorLedgerNotExist, c)
		return
	}
	//更新账本
	if err := db.Model(&ledger2.Ledger{}).Where("id = ?", ledger.ID).Updates(&ledger).Error; err != nil {
		code.FailResponse(code.ErrorUpdateLedger, c)
		return
	}
	code.SuccessResponse(ledger, code.SUCCESS, c)
}

// GetLedgerList 获取账本列表
//@Summary 获取账本列表
//@Tags ledger
//@Accept  json
//@Produce  json
//@Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.Ledger}
//@Router /ledger/list [get]
func (*ApiLedger) GetLedgerList(c *gin.Context) {
	userId := utils.FindUserID(c)
	data, cd, err := ledgerService.GetLedgerList(userId)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// GetLedgerDetail 获取账本详情
//@Summary 获取账本详情
//@Tags ledger
//@Accept  json
//@Produce  json
//@Param id query int true "账本ID"
//@Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.Ledger}
//@Router /ledger/detail [get]
func (*ApiLedger) GetLedgerDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := ledgerService.GetLedgerDetail(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
