package admin

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
)

type UserAdminApi struct {
}

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags admin/user
// @Accept  json
// @Produce  json
// @Param  query  query    system.SysUserRequest  true  "用户查询参数"
// @Success 200 {object} code.Response{data=system.User,code=int,total=int64,msg=string,success=bool}
// @Router /admin/user/list [get]
func (*UserAdminApi) GetUserList(c *gin.Context) {
	var query system.SysUserRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	data, total, cd, err := userAdminService.GetUserList(query, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}

// AddRole 新增角色
// @Summary 新增角色
// @Description 新增角色
// @Tags admin/role
// @Accept  json
// @Produce  json
// @Param  role  body  system.Role  true  "角色"
// @Success 200 {object} code.Response{data=string,code=int,msg=string,success=bool}
// @Router /admin/user/add/role [post]
func (*UserAdminApi) AddRole(c *gin.Context) {
	var role system.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	if err := userAdminService.AddRole(role); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	code.SuccessResponse(nil, code.SUCCESS, c)
}

// 角色列表 GetRoleList
// @Summary 角色列表
// @Description 角色列表
// @Tags admin/role
// @Accept  json
// @Produce  json
// @Param  query  query    string  true  "角色名称"
// @Param  query  query    models.PaginationRequest  true  "分页"
// @Success 200 {object} code.Response
// @Router /admin/user/role/list [get]
func (*UserAdminApi) GetRoleList(c *gin.Context) {
	data, total, cd, err := userAdminService.GetRoleList(c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}

// UpdateRole 修改角色
// @Summary 修改角色
// @Description 修改角色
// @Tags admin/role
// @Accept  json
// @Produce  json
// @Param  role  body  system.Role  true  "角色"
// @Success 200 {object} code.Response
// @Router /admin/user/update/role [put]
func (*UserAdminApi) UpdateRole(c *gin.Context) {
	var role system.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	if cd, err := userAdminService.UpdateRole(role); err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, code.SUCCESS, c)
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Description 删除角色
// @Tags admin/role
// @Accept  json
// @Produce  json
// @Param  id  query  int  true  "角色ID"
// @Success 200 {object} code.Response
// @Router /admin/user/delete/role [post]
func (*UserAdminApi) DeleteRole(c *gin.Context) {
	id := c.Query("id")
	if cd, err := userAdminService.DeleteRole(id); err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, code.SUCCESS, c)
}
