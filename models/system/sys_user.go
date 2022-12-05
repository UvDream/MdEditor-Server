package system

import (
	"gorm.io/gorm"
	"server/models"
	"server/utils"
)

// User 用户表
type User struct {
	models.Model
	UserName string `json:"user_name" gorm:"comment:用户名"`
	NickName string `json:"nick_name" gorm:"comment:昵称"`
	Password string `json:"password" gorm:"comment:密码"`
	Phone    string `json:"phone" gorm:"comment:手机号"`
	Email    string `json:"email" gorm:"comment:邮箱"`
	Avatar   string `json:"avatar" gorm:"comment:头像"`
	//关联到角色表
	Roles        []SysRole  `json:"roles" gorm:"many2many:sys_user_role;"`
	UserConfigID string     `json:"user_config_id" gorm:"comment:用户配置ID"`
	UserConfig   UserConfig `json:"user_config" gorm:"foreignKey:UserConfigID"`
	//	来源
	Source string `json:"source" gorm:"comment:来源;default:'editor'"`
	//	微信用户唯一ID
	OpenID string `json:"open_id" gorm:"comment:微信用户唯一ID"`
}

//LoginRequest 登陆请求参数
type LoginRequest struct {
	Username  string `json:"user_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}
type RetrievePasswordRequest struct {
	UserName string `json:"user_name" gorm:"comment:用户名"`
	NickName string `json:"nick_name" gorm:"comment:昵称"`
	Password string `json:"password" gorm:"comment:密码"`
	Phone    string `json:"phone" gorm:"comment:手机号"`
	Email    string `json:"email" gorm:"comment:邮箱"`
}

// SysUserRequest 用户列表请求参数
type SysUserRequest struct {
	models.PaginationRequest
	Username string `form:"username" json:"user_name"`
	Nickname string `form:"nickname" json:"nick_name"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.Password != "" {
		user.Password = utils.BcryptHash(user.Password)
	}
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if user.Password == "" {
		tx.Statement.Omit("password")
	}
	return
}
