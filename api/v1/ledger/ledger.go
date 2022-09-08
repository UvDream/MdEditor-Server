package ledger

import (
	"github.com/gin-gonic/gin"
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
}

// GetLedgerList 获取账本列表
//@Summary 获取账本列表
//@Tags ledger
//@Accept  json
//@Produce  json
//@Param    query  query    models.PaginationRequest  false  "参数"
//@Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]ledger.Ledger}
//@Router /ledger/list [get]
func (*ApiLedger) GetLedgerList(c *gin.Context) {
}
