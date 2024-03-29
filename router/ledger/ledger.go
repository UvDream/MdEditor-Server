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
		ledgerRouter.GET("/detail", accountApi.GetLedgerDetail)
		//共享账本,生成邀请码
		ledgerRouter.POST("/share", accountApi.ShareLedger)
		//加入账本
		ledgerRouter.POST("/join", accountApi.JoinLedger)
		//邀请加入账本
		ledgerRouter.POST("/invite", accountApi.InviteLedger)
		//分类
		ledgerRouter.POST("/category/create", accountApi.CreateLedgerCategory)
		ledgerRouter.DELETE("/category/delete", accountApi.DeleteLedgerCategory)
		ledgerRouter.PUT("/category/update", accountApi.UpdateLedgerCategory)
		ledgerRouter.GET("/category/list", accountApi.GetLedgerCategoryList)
		ledgerRouter.GET("/category/detail", accountApi.GetLedgerCategoryDetail)
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
		ledgerRouter.GET("/bill/detail", accountApi.GetBillDetail)
		ledgerRouter.GET("/bill/normal/list", accountApi.GetBillNormalList)
		//获取图标列表
		ledgerRouter.GET("/icon/list", accountApi.GetIconList)
		//获取图标颜色列表
		ledgerRouter.GET("/icon/color/list", accountApi.GetIconColorList)
		//根据分类ID获取账单列表
		ledgerRouter.GET("/bill/category/list", accountApi.GetBillListByCategoryId)
		//预算
		ledgerRouter.POST("/budget/create", accountApi.CreateBudget)
		ledgerRouter.DELETE("/budget/delete", accountApi.DeleteBudget)
		ledgerRouter.PUT("/budget/update", accountApi.UpdateBudget)
		ledgerRouter.GET("/budget/list", accountApi.GetBudgetList)
		//删除预算
		ledgerRouter.DELETE("/budget/batch_deletion", accountApi.BatchDeletion)
		//日历统计
		ledgerRouter.GET("/statistics/calendar", accountApi.GetCalendar)
		//首页统计
		ledgerRouter.GET("/statistics/home", accountApi.GetHomeStatistics)
		//分类统计
		ledgerRouter.GET("/statistics/category", accountApi.GetCategoryStatistics)
		//支出/收入统计
		ledgerRouter.GET("/statistics/income_expenditure", accountApi.GetIncomeExpenditureStatistics)
		//成员统计
		ledgerRouter.GET("/statistics/member", accountApi.GetMemberStatistics)
		ledgerRouter.GET("/statistics/personal", accountApi.GetPersonalStatistics)
		//分类详细统计
		ledgerRouter.GET("/statistics/category_detail", accountApi.GetCategoryDetailStatistics)
		//金额统计
		ledgerRouter.GET("/statistics/amount", accountApi.GetAmountStatistics)
		//App相关
		//是否有更新
		ledgerRouter.GET("/app/need_update", accountApi.GetAppNeedUpdate)
		//周期记账
		//新增
		ledgerRouter.POST("/loop_account", accountApi.LoopAccount)
		//修改
		ledgerRouter.PUT("/loop_account", accountApi.UpdateLoopAccount)
		//删除周期记账
		ledgerRouter.DELETE("/loop_account/delete", accountApi.DeleteLoopAccount)
		//改变周期记账状态
		ledgerRouter.PUT("/loop_account/change_status", accountApi.ChangeLoopAccountStatus)
	}
	return ledgerRouter
}
