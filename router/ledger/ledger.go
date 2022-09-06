package ledger

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type AccountsRouter struct{}

func (*AccountsRouter) InitLedgerRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	ledgerRouter := Router.Group("ledger")
	accountApi := v1.ApiGroupApp.LedgerApiGroup
	{
		ledgerRouter.POST("/create", accountApi.CreateLedger)
		ledgerRouter.DELETE("/delete", accountApi.DeleteLedger)
		ledgerRouter.PUT("/update", accountApi.UpdateLedger)
		ledgerRouter.GET("/list", accountApi.GetLedgerList)
		ledgerRouter.GET("/detail", accountApi.GetLedgerDetail)
	}
	return ledgerRouter
}
