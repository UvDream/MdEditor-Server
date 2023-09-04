package admin

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type LedgerAdminStruct struct {
}

func (*LedgerAdminStruct) InitLedgerAdminRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	adminRouter := Router.Group("ledger")
	ledgerAdminApi := v1.ApiGroupApp.AdminAPiGroup.LedgerAdminApi
	{
		//	账本列表
		adminRouter.GET("/list", ledgerAdminApi.GetLedgerList)
		//	账单列表
		adminRouter.GET("/bill/list", ledgerAdminApi.GetBillList)
		//	颜色列表
		adminRouter.GET("/color/list", ledgerAdminApi.GetColorList)
		//	新增颜色
		adminRouter.POST("/color/add", ledgerAdminApi.AddColor)
		//	删除颜色
		adminRouter.DELETE("/color/delete", ledgerAdminApi.DeleteColor)
		//	icon列表
		adminRouter.GET("/icon/list", ledgerAdminApi.GetIconList)
		//	新增icon
		adminRouter.POST("/icon/add", ledgerAdminApi.AddIcon)
		//	删除icon
		adminRouter.DELETE("/icon/delete", ledgerAdminApi.DeleteIcon)
	}
	return adminRouter
}
