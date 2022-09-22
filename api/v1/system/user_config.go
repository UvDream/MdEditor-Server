package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
	"server/utils"
)

//SetUserConfig 设置用户配置
// @Tags system
// @Summary 设置用户配置
// @Produce  json
// @Param  config body system.UserConfig true "用户配置"
// @Success 200 {string} code.Response {"success":true,"data":system.UserConfig,"msg":"设置成功"}
// @Router /user/user_config [post]
func (*UserApi) SetUserConfig(c *gin.Context) {
	var userConfig system.UserConfig
	if err := c.ShouldBindJSON(&userConfig); err != nil {
		code.FailResponse(code.ErrorSetUserConfigMissingParam, c)
		return
	}
	userID := utils.FindUserID(c)
	data, cd, err := userService.SetUserConfig(userConfig, userID)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

//EditUserConfig 设置用户配置
// @Tags system
// @Summary 修改用户配置
// @Produce  json
// @Param  config body system.UserConfig true "用户配置"
// @Success 200 {string} code.Response {"success":true,"data":system.UserConfig,"msg":"设置成功"}
// @Router /user/user_config [put]
func (*UserApi) EditUserConfig(c *gin.Context) {
	var userConfig system.UserConfig
	if err := c.ShouldBindJSON(&userConfig); err != nil {
		code.FailResponse(code.ErrorSetUserConfigMissingParam, c)
		return
	}
	if userConfig.ID == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	data, cd, err := userService.EditorUserConfig(userConfig)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// GetUserConfig 获取用户配置
// @Tags system
// @Summary 获取用户配置
// @Produce  json
// @Success 200 {object} code.Response {"success":true,"data":system.UserConfig,"msg":"设置成功"}
// @Router /user/user_config [get]
func (*UserApi) GetUserConfig(c *gin.Context) {
	userID := utils.FindUserID(c)
	data, cd, err := userService.GetUserConfig(userID)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
