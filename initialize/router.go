package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/middleware"
	"server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	InitSwag(Router)
	//Router.Static("/files", "./files")
	//设置静态文件夹
	Router.StaticFS("/files", gin.Dir(global.Config.Local.Path, true))
	//系统相关路由
	systemRouter := router.RoutersGroupApp.System
	//文章相关路由
	articleRouter := router.RoutersGroupApp.Article
	//标签相关路由
	tagRouter := router.RoutersGroupApp.Tag
	//分类相关路由
	categoryRouter := router.RoutersGroupApp.Category
	//文件相关路由
	fileRouter := router.RoutersGroupApp.File
	//主题相关路由
	themeRouter := router.RoutersGroupApp.Theme
	//记账相关路由
	ledgerRouter := router.RoutersGroupApp.Account
	//halo相关路由
	haloRouter := router.RoutersGroupApp.Halo
	adminRouter := router.RoutersGroupApp.Admin
	//跨域设置
	Router.Use(middleware.Cors()) //放行所有的请求
	//Router.Use(middleware.CorsByRules())  //按照配置规则放行跨域
	//公开的接口
	PublicGroup := Router.Group("/public")
	{
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "程序正常运行",
			})
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}
	//web端路由
	PrivateGroup := Router.Group("/")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		//systemRouter.InitSysRouter(PrivateGroup)
		articleRouter.InitArticleRouter(PrivateGroup)
		tagRouter.InitTagRouter(PrivateGroup)
		categoryRouter.InitCategoriesRouter(PrivateGroup)
		fileRouter.InitFileRouter(PrivateGroup)
		themeRouter.InitThemeRouter(PrivateGroup)
		ledgerRouter.InitLedgerRouter(PrivateGroup)
		haloRouter.InitHaloRouter(PrivateGroup)
	}
	//管理端路由
	AdminGroup := Router.Group("/admin")
	AdminGroup.Use(middleware.JWTAuth())
	{
		adminRouter.InitUserAdminRouter(AdminGroup)
		adminRouter.InitLedgerAdminRouter(AdminGroup)
	}
	return Router
}
