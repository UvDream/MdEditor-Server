package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("/base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("/login", baseApi.Login)
		baseRouter.POST("/register", baseApi.Register)
		baseRouter.GET("/init_data", baseApi.InitData)
		baseRouter.POST("/captcha", baseApi.Captcha)
		baseRouter.GET("/md", baseApi.GetArticleMd)
		//找回密码
		baseRouter.POST("/retrieve_password", baseApi.RetrievePassword)
		//获取open_id
		baseRouter.GET("/wx/get_openid", baseApi.GetOpenId)
		//利用用户信息生成token
		baseRouter.POST("/wx/get_token", baseApi.GetToken)
		//发送邮件验证码
		baseRouter.GET("/send_email_code", baseApi.SendEmailCode)
		//验证邮件验证是否正确
		baseRouter.GET("/verify_email_code", baseApi.VerifyEmailCode)
	}
	return baseRouter
}
