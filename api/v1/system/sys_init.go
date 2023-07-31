package system

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/system"
)

// InitData 初始化数据
// @Summary 初始化数据
// @Description 初始化数据
// @Tags system
// @ID InitData
// @Accept  json
// @Produce  json
// @Success 200 {object} code.Response"{"code":200,"data":{},"msg":"ok"}"
// @Router /public/base/init_data [get]
func (b *BaseApi) InitData(c *gin.Context) {
	//查询是否存在admin
	var user system.User
	global.DB.Where("user_name = ?", "admin").First(&user)
	//不存在则初始化
	if user.ID == "" {
		user = system.User{
			UserName: "admin",
			Password: "qwer1314520.",
			NickName: "admin",
			Avatar:   "https://pic.imgdb.cn/item/64c0cc451ddac507ccd49532.png",
			Phone:    "17621953630",
			Email:    "uvdream@163.com",
		}
		global.DB.Create(&user)
	}
	//查询是否存在角色
	var role system.Role
	global.DB.Where("role_name = ?", "admin").First(&role)
	//不存在则初始化
	if role.ID == "" {
		role = system.Role{
			RoleName:  "超级管理员",
			Remark:    "超级管理员",
			RoleKey:   "root",
			IsDefault: "1",
		}
		global.DB.Create(&role)
	}
	//初始化角色和admin关联
	var userRole system.UserRole
	global.DB.Where("user_id = ?", user.ID).First(&userRole)
	if userRole.UserID == "" {
		userRole = system.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}
		global.DB.Create(&userRole)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "初始化数据成功",
	})
}
