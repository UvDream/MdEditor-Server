package halo

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/service"
)

type ApiHaloGroup struct {
}

var serviceHalo = service.ServicesGroupApp.HaloServiceGroup

type haloUser struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// GetToken 获取token
func (*ApiHaloGroup) GetToken(c *gin.Context) {
	var query haloUser
	if err := c.ShouldBindQuery(&query); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	serviceHalo.GetToken()
}
