package ledger

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
	"server/utils"
)

// LoopAccount 周期记账
// @Tags Ledger/loop
// @Summary 周期记账
// @Description 周期记账
// @Accept  json
// @Produce  json
// @Param loop_account body ledger.LoopAccount true "周期记账"
// @Success 200 {object}  code.Response{}
// @Router /ledger/loop_account [post]
func (*ApiLedger) LoopAccount(c *gin.Context) {
	var loopAccount ledger.LoopAccount
	err := c.ShouldBindJSON(&loopAccount)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	loopAccount.CreatorID = utils.FindUserID(c)
	cd, err := ledgerService.LoopAccountService(loopAccount)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// UpdateLoopAccount 更新周期记账
// @Tags Ledger/loop
// @Summary 更新周期记账
// @Description 更新周期记账
// @Accept  json
// @Produce  json
// @Param loop_account body ledger.LoopAccount true "周期记账"
// @Success 200 {object}  code.Response{}
// @Router /ledger/loop_account [put]
func (*ApiLedger) UpdateLoopAccount(c *gin.Context) {
	var loopAccount ledger.LoopAccount
	err := c.ShouldBindJSON(&loopAccount)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	cd, err := ledgerService.UpdateLoopAccountService(loopAccount)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// DeleteLoopAccount 删除周期记账
// @Tags Ledger/loop
// @Summary 删除周期记账
// @Description 删除周期记账
// @Accept  json
// @Produce  json
// @Param id query string true "周期记账ID"
// @Success 200 {object}  code.Response{}
// @Router /ledger/loop_account/delete [delete]
func (*ApiLedger) DeleteLoopAccount(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerService.DeleteLoopAccountService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// ChangeLoopAccountStatus 改变周期记账状态
// @Tags Ledger/loop
// @Summary 改变周期记账状态
// @Description 改变周期记账状态
// @Accept  json
// @Produce  json
// @Param id query string true "周期记账ID"
// @Success 200 {object}  code.Response{}
// @Router /ledger/loop_account/change_status [put]
func (*ApiLedger) ChangeLoopAccountStatus(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerService.ChangeLoopAccountStatusService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}
