package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
)

type BaseApi struct{}

// Login 登录
// @Tags user
// @Summary 用户登录
// @Produce  application/json
// @Param data body system.LoginRequest true "用户名, 密码, 验证码"
// @Success 200 {object} code.Response{data=system.User,code=int,msg=string,success=bool}
// @Router  /public/base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var loginRequest system.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	//首先验证参数是否齐全
	if err != nil {

		return
	}
	//其次验证验证码是否正确
	//if err := b.VerifyCaptcha(loginRequest.Captcha, loginRequest.CaptchaId); err != true {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "验证码错误",
	//	})
	//	return
	//}
	//最后验证用户名和密码是否正确
	if user, token, msg, err := userService.Login(loginRequest.Username, loginRequest.Password); err != nil {
		code.FailWithMessage(msg, c)
		return
	} else {
		code.OkWithDetailed(gin.H{"token": token, "user_info": user}, msg, c)
	}
}

// Register 注册
// @Tags user
// @Summary 用户注册
// @Produce  application/json
// @Param data body system.User true "用户名, 密码, 验证码"
// @Success 200 {object} code.Response{data=system.User,code=int,msg=string,success=bool}
// @Router  /public/base/register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var registerRequest system.User
	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		code.FailResponse(code.ErrorRegisterMissingParam, c)
		return
	}
	data, cd, err := userService.RegisterService(registerRequest)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

//GetArticleDetail 查询文章详情
//@Summary 查询文章详情
//@Tags article
//@Accept  json
//@Produce  json
//@Param        id   query     string  true  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":article.Article,"msg":"操作成功"}"
//@Router /public/base/detail [get]
func (*BaseApi) GetArticleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailWithMessage("id不能为空", c)
		return
	}
	articleContent, msg, err := articleService.GetArticleDetailService(id)
	if err != nil {
		code.FailWithMessage(msg, c)
		return
	}
	code.OkWithDetailed(articleContent, msg, c)
}
