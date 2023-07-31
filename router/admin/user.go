package admin

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserAdminStruct struct {
}

func (*UserAdminStruct) InitUserAdminRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	adminRouter := Router.Group("user")
	userAdminApi := v1.ApiGroupApp.AdminAPiGroup.UserAdminApi
	{
		//用户列表
		adminRouter.GET("/list", userAdminApi.GetUserList)
		//	新增角色
		adminRouter.POST("/add/role", userAdminApi.AddRole)
		//	角色列表
		adminRouter.GET("/role/list", userAdminApi.GetRoleList)
	}
	return adminRouter
}
