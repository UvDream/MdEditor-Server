package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
	"server/utils"
)

type UserApi struct{}

func (b *UserApi) UserList(c *gin.Context) {
	var userRequest system.SysUserRequest
	//先判断参数是否合法
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		code.FailWithMessage(err.Error(), c)
	}
	userList, total, msg, err := userService.GetUserListService(&userRequest)
	if err != nil {
		code.FailWithMessage(msg, c)
	}
	code.OkWithDetailed(gin.H{
		"list":  userList,
		"total": total,
	}, msg, c)
}

// GetUserInfo 获取用户信息
// @Tags user
// @Summary 获取用户信息
// @Produce  application/json
// @Success 200 {object} code.Response{data=system.User,code=int,msg=string,success=bool}
// @Router  /user/get_user_info [get]
func (*UserApi) GetUserInfo(c *gin.Context) {
	userID := utils.FindUserID(c)
	user, cd, err := userService.GetUserInfoService(userID)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(user, cd, c)
}

//UnbindEmail 解绑邮箱
// @Tags user
// @Summary 解绑邮箱
// @Produce  application/json
// @Success 200 {object} code.Response{data=string,code=int,msg=string,success=bool}
// @Router  /user/unbind_email [post]
func (*UserApi) UnbindEmail(c *gin.Context) {
	userID := utils.FindUserID(c)
	cd, err := userService.UnbindEmailService(userID)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse("", cd, c)
}

// BindEmail 绑定邮箱
// @Tags user
// @Summary 绑定邮箱
// @Accept  json
// @Produce  json
// @Param article body system.BindEmail true "创建文章"
// @Success 200 {object} code.Response{data=string,code=int,msg=string,success=bool}
// @Router  /user/bind_email [post]
func (*UserApi) BindEmail(c *gin.Context) {
	userID := utils.FindUserID(c)
	var bindEmail system.BindEmail
	if err := c.ShouldBindJSON(&bindEmail); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	cd, err := userService.BindEmailService(c, userID, bindEmail)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse("绑定成功", cd, c)
}
