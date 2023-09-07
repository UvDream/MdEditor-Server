package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
	"server/utils"
)

// MenuList 菜单列表
// @Tags system/role
// @Summary 菜单列表
// @Produce  json
// @Success 200 {object} code.Response{success=bool,data=[]string,msg=string}
// @Router /user/menu_list [get]
func (*UserApi) MenuList(c *gin.Context) {
	data, cd, err := userService.MenuListService()
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// AddMenu 新增菜单
// @Tags system/role
// @Summary 新增菜单
// @Produce  json
// @Param  menu body system.Permission true "菜单信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/add_menu [post]
func (*UserApi) AddMenu(c *gin.Context) {
	data := system.Permission{}
	if err := c.ShouldBindJSON(&data); err != nil {
		code.FailResponse(code.ErrorAddMenuFail, c)
		return
	}
	cd, err := userService.AddMenuService(data)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// DeleteMenu 删除菜单
// @Tags system/role
// @Summary 删除菜单
// @Produce  json
// @Param  id query string true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/delete_menu [delete]
func (*UserApi) DeleteMenu(c *gin.Context) {
	id := c.Query("id")
	cd, err := userService.DeleteMenuService(id)
	if err != nil {
		code.FailResponse(code.ErrorDeleteMenuFail, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// UserRole 分配用户角色
// @Tags system/role
// @Summary 分配用户角色
// @Produce  json
// @Param  query body system.UserRoleRequest true "分配用户角色"
// @Success 200 {object} code.Response "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/user_role [post]
func (*UserApi) UserRole(c *gin.Context) {
	var query system.UserRoleRequest
	if err := c.ShouldBindJSON(&query); err != nil {
		code.FailResponse(code.ErrorUserRoleFail, c)
		return
	}
	cd, err := userService.UserRoleService(query)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// GetPermission 获取权限
// @Tags system/role
// @Summary 获取权限
// @Produce  json
// @Success 200 {object} code.Response
// @Router /user/get_permission [get]
func (*UserApi) GetPermission(c *gin.Context) {
	userId := utils.FindUserID(c)
	data, cd, err := userService.GetPermissionService(userId)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}
