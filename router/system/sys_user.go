package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		//	获取用户信息
		userRouter.GET("get_user_info", baseApi.GetUserInfo)
		//	解绑邮箱
		userRouter.POST("unbind_email", baseApi.UnbindEmail)
		//	绑定邮箱
		userRouter.POST("bind_email", baseApi.BindEmail)
		userRouter.POST("user_list", baseApi.UserList)
		//	配置
		//增
		userRouter.POST("user_config", baseApi.SetUserConfig)
		//删
		//改
		userRouter.PUT("user_config", baseApi.EditUserConfig)
		//查
		userRouter.GET("user_config", baseApi.GetUserConfig)
		//	生成用户邀请码
		userRouter.POST("user_invite_code", baseApi.SetUserInviteCode)
		//	获取用户邀请码
		userRouter.GET("user_invite_code", baseApi.GetUserInviteCode)
		//	获取用户邀请码列表
		userRouter.GET("user_invite_code_list", baseApi.GetUserInviteCodeList)
		//	填写邀请码
		userRouter.POST("fill_user_invite_code", baseApi.FillUserInviteCode)
		//	建议反馈
		userRouter.POST("feedback", baseApi.Feedback)
		// 获取建议反馈列表
		userRouter.GET("feedback_list", baseApi.FeedbackList)
		//	查询是否是会员
		userRouter.GET("is_vip", baseApi.IsVip)

	}
	return userRouter
}
