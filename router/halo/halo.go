package halo

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type RouteHaloGroup struct{}

func (*RouteHaloGroup) InitHaloRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	haloRouter := Router.Group("halo")
	haloApi := v1.ApiGroupApp.HaloApiGroup
	{
		haloRouter.GET("/token", haloApi.GetToken)
		haloRouter.POST("/save", haloApi.SaveArticle)
		//	tags
		haloRouter.GET("/tags", haloApi.GetTags)
		//	category
		haloRouter.GET("/category", haloApi.GetCategory)
	}
	return haloRouter
}
