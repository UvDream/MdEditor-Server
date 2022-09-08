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
		//账本
		ledgerRouter.POST("/create", accountApi.CreateLedger)
		ledgerRouter.DELETE("/delete", accountApi.DeleteLedger)
		ledgerRouter.PUT("/update", accountApi.UpdateLedger)
		ledgerRouter.GET("/list", accountApi.GetLedgerList)
		//分类
		ledgerRouter.POST("/category/create", accountApi.CreateLedgerCategory)
		ledgerRouter.DELETE("/category/delete", accountApi.DeleteLedgerCategory)
		ledgerRouter.PUT("/category/update", accountApi.UpdateLedgerCategory)
		ledgerRouter.GET("/category/list", accountApi.GetLedgerCategoryList)
		//标签
		ledgerRouter.POST("/tag/create", accountApi.CreateLedgerTag)
		ledgerRouter.DELETE("/tag/delete", accountApi.DeleteLedgerTag)
		ledgerRouter.PUT("/tag/update", accountApi.UpdateLedgerTag)
		ledgerRouter.GET("/tag/list", accountApi.GetLedgerTagList)
		//账单
		ledgerRouter.POST("/bill/create", accountApi.CreateBill)
		ledgerRouter.DELETE("/bill/delete", accountApi.DeleteBill)
		ledgerRouter.PUT("/bill/update", accountApi.UpdateBill)
		ledgerRouter.GET("/bill/list", accountApi.GetBillList)
	}
	return ledgerRouter
}
