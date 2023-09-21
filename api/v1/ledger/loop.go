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
