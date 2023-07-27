package system

import (
	"gorm.io/gorm"
	"server/models"
	"server/utils"
	"time"
)

// User 用户表
type User struct {
	models.Model
	UserName string `json:"user_name" gorm:"comment:用户名"`
	NickName string `json:"nick_name" gorm:"comment:昵称"`
	Password string `json:"password" gorm:"comment:密码"`
	// 个人描述
	Desc   string `json:"desc" gorm:"comment:个人描述"`
	Phone  string `json:"phone" gorm:"comment:手机号"`
	Email  string `json:"email" gorm:"comment:邮箱"`
	Avatar string `json:"avatar" gorm:"comment:头像",default:"https://pic.imgdb.cn/item/64c0cc451ddac507ccd49532.png"`
	//性别
	Gender string `json:"gender" gorm:"comment:性别;default:'1'"` //1 男 2 女 3 保密
	//关联到角色表
	Roles        []SysRole  `json:"roles" gorm:"many2many:sys_user_role;"`
	UserConfigID string     `json:"user_config_id" gorm:"comment:用户配置ID"`
	UserConfig   UserConfig `json:"user_config" gorm:"foreignKey:UserConfigID"`
	//	来源
	Source string `json:"source" gorm:"comment:来源;default:'editor'"`
	//	微信用户唯一ID
	OpenID string `json:"open_id" gorm:"comment:微信用户唯一ID"`
	//	邀请码
	InviteCode string `json:"invite_code" gorm:"comment:邀请码"`
	//	会员信息
	MemberID string `json:"member_id" gorm:"comment:会员信息"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID"`
	// 被邀请码
	InvitedCode string `json:"invited_code" gorm:"comment:被邀请码"`
}

// Member 会员信息
type Member struct {
	models.Model
	//  是否是会员
	IsMember bool `json:"is_member" gorm:"comment:是否是会员,default:false"`
	//	会员到期时间
	ExpireTime time.Time `json:"expire_time" gorm:"comment:会员到期时间"`
	//	会员状态
	Status int `json:"status" gorm:"comment:会员状态"`
}

// LoginRequest 登陆请求参数
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
type BindEmail struct {
	Email string `json:"email" binding:"required"`
	//	唯一ID
	UniqueID string `json:"unique_id" binding:"required"`
	//	验证码
	Code string `json:"code" binding:"required"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.Password == "" {
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
