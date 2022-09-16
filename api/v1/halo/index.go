package halo

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/halo"
	"server/service"
)

type ApiHaloGroup struct {
}

var serviceHalo = service.ServicesGroupApp.HaloServiceGroup

// GetToken 获取token
// @Summary 获取token
// @Tags halo
// @Accept  json
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param url query string true "url"
// @Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
// @Router /halo/token [get]
func (*ApiHaloGroup) GetToken(c *gin.Context) {
	var query halo.UserHalo
	if err := c.ShouldBindQuery(&query); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	token, cd, err := serviceHalo.GetToken(query)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(token, cd, c)
}

// SaveArticle 保存文章
// @Summary 保存文章
// @Tags halo
// @Accept  json
// @Produce  json
//@Param article body halo.ArticleHaloResponse true "创建文章"
// @Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
// @Router /halo/save [post]
func (*ApiHaloGroup) SaveArticle(c *gin.Context) {
	var query halo.ArticleHaloResponse
	if err := c.ShouldBindJSON(&query); err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	token, cd, err := serviceHalo.SaveArticle(query)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(token, cd, c)
}

// GetTags 获取标签
// @Summary 获取标签
// @Tags halo
// @Accept  json
// @Produce  json
// @Param token query string true "token"
// @Param url query string true "url"
// @Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
// @Router /halo/tags [get]
func (*ApiHaloGroup) GetTags(c *gin.Context) {
	token := c.Query("token")
	url := c.Query("url")
	if token == "" {
		code.FailWithMessage("token不能为空", c)
		return
	}
	tags, cd, err := serviceHalo.GetTags(url, token)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(tags, cd, c)
}

//GetCategory 获取分类
// @Summary 获取分类
// @Tags halo
// @Accept  json
// @Produce  json
// @Param token query string true "token"
// @Param url query string true "url"
// @Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
// @Router /halo/category [get]
func (*ApiHaloGroup) GetCategory(c *gin.Context) {
	token := c.Query("token")
	url := c.Query("url")
	if token == "" {
		code.FailWithMessage("token不能为空", c)
		return
	}
	category, cd, err := serviceHalo.GetCategory(url, token)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(category, cd, c)
}
