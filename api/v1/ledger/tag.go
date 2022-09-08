package ledger

import "github.com/gin-gonic/gin"

// CreateLedgerTag 创建账本标签
// @Summary 创建账本标签
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.LedgerTag true "创建账单标签"
// @Success 200 {object} code.Response{data=ledger.LedgerTag,code=int,msg=string,success=bool}
// @Router /ledger/tag/create [post]
func (*ApiLedger) CreateLedgerTag(c *gin.Context) {
}

// DeleteLedgerTag 删除账本标签
// @Summary 删除账本标签
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param id query int true "账本标签ID"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool}
// @Router /ledger/tag/delete [delete]
func (*ApiLedger) DeleteLedgerTag(c *gin.Context) {

}

// UpdateLedgerTag 更新账本标签
// @Summary 更新账本标签
// @Tags ledger
// @Accept  json
// @Produce  json
// @Param article body ledger.LedgerTag true "更新账单标签"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=ledger.LedgerTag}
// @Router /ledger/tag/update [put]
func (*ApiLedger) UpdateLedgerTag(c *gin.Context) {

}

// GetLedgerTagList 获取账本标签列表
// @Summary 获取账本标签列表
// @Tags ledger
// @Accept  json
// @Produce  json
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.LedgerTag}
// @Router /ledger/tag/list [get]
func (*ApiLedger) GetLedgerTagList(c *gin.Context) {

}
