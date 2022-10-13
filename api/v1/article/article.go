package article

import (
	"github.com/gin-gonic/gin"
	"server/code"
	article2 "server/models/article"
	"server/utils"
)

type ArticlesApi struct{}

// CreateArticle 创建文章
//@Summary 创建文章
//@Tags article
//@Accept  json
//@Produce  json
//@Param article body article2.Article true "创建文章"
//@Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
//@Router /article/create [post]
func (article *ArticlesApi) CreateArticle(c *gin.Context) {
	var articleOpts article2.Article
	//校验必填信息
	err := c.ShouldBindJSON(&articleOpts)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	articleOpts.UserID = utils.FindUserID(c)
	articleContent, msg, err := articleService.CreateArticle(articleOpts)
	if err != nil {
		code.FailWithMessage(msg, c)
		return
	}
	code.OkWithDetailed(articleContent, msg, c)
}

// DeleteArticle 删除文章
//@Summary 删除文章
//@Tags article
//@Accept  json
//@Produce  json
//@Param        id   query     string  true  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
//@Router /article/delete [delete]
func (article *ArticlesApi) DeleteArticle(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailWithMessage("id不能为空", c)
	}
	msg, err := articleService.DeleteArticleService(id)
	if err != nil {
		code.FailWithMessage(msg, c)
		return
	}
	code.OkWithMessage(msg, c)
}

// UpdateArticle 修改文章
//@Summary 修改文章
//@Tags article
//@Accept  json
//@Produce  json
//@Param article body article2.Article true "修改文章"
//@Success 200 {object} code.Response "{"code":200,"data":article2.Article,"msg":"操作成功"}"
//@Router /article/update [put]
func (article *ArticlesApi) UpdateArticle(c *gin.Context) {
	var articleOpts article2.Article
	//校验必填信息
	err := c.ShouldBindJSON(&articleOpts)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	if articleOpts.ID == "" {
		code.FailWithMessage("缺少必要参数uuid", c)
		return
	}
	articleContent, msg, err := articleService.UpdateArticleService(articleOpts)
	if err != nil {
		code.FailWithMessage(msg, c)
		return
	}
	code.OkWithDetailed(articleContent, msg, c)

}

// GetArticleList 查询文章
//@Summary 查询文章
//@Tags article
//@Accept  json
//@Produce  json
//@Param 	query query article2.ListArticleRequest true "查询文章"
//@Param    query  query    models.PaginationRequest  false  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":[]article.Article,"msg":"操作成功"}"
//@Router /article/list [get]
func (article *ArticlesApi) GetArticleList(c *gin.Context) {
	//x-token
	xToken := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(xToken)
	if err != nil {
		code.FailWithMessage("token验证失败", c)
		return
	}
	var articleOpts article2.ListArticleRequest
	err = c.ShouldBindQuery(&articleOpts)
	if err != nil {
		code.FailResponse(code.ErrorGetQueryParam, c)
		return
	}
	list, total, msg, err := articleService.GetArticleListService(articleOpts, claims.ID, c)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	code.OkWithDetailed(gin.H{
		"list":  list,
		"total": total,
	}, msg, c)
}

//GetArticleHistory 查询文章历史记录
//@Summary 查询文章历史记录
//@Tags article
//@Accept  json
//@Produce  json
//@Param        id   query     string  true  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":[]article.Article,"msg":"操作成功"}"
//@Router /article/history [get]
func (article *ArticlesApi) GetArticleHistory(c *gin.Context) {
	uid := c.Query("id")
	if uid == "" {
		code.FailWithMessage("id不能为空", c)
		return
	}
	list, msg, err := articleService.GetArticleHistoryService(uid)
	if err != nil {
		code.FailWithMessage(msg, c)
		return
	}
	code.OkWithDetailed(gin.H{
		"list": list,
	}, "查询成功", c)
}

//GetArticleDetail 查询文章详情
//@Summary 查询文章详情
//@Tags article
//@Accept  json
//@Produce  json
//@Param        id   query     string  true  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":article.Article,"msg":"操作成功"}"
//@Router /article/detail [get]
func (*ArticlesApi) GetArticleDetail(c *gin.Context) {
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

// ExportMd 文章导出md
//@Summary 文章导出md
//@Tags article
//@Accept  json
//@Produce  json
//@Param        id   query     string  true  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":{},"msg":"操作成功"}"
//@Router /article/export [get]
func (*ArticlesApi) ExportMd(c *gin.Context) {
	userID := utils.FindUserID(c)
	articleID := c.Query("id")
	if articleID == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	path, cd, err := articleService.ExportMd(userID, articleID)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(path, cd, c)

}
