package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
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
