package article

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/article"
)

type CategoriesApi struct{}

//CreateCategory 创建Category
//@Summary 创建Category
//@Description 创建Category
//@Tags category
//@Accept  json
//@Produce  json
//@Param category body article.Category true "创建Category"
//@Success 200 {object} code.Response"{"code":200,"data":article.Category,"msg":"操作成功"}"
//@Router /category/create [post]
func (ct *CategoriesApi) CreateCategory(c *gin.Context) {
	var category article.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	if list, msg, err := categoryService.CreateCategoryService(category); err != nil {
		code.FailWithMessage(msg, c)
		return
	} else {
		code.OkWithDetailed(list, msg, c)
	}
}

//DeleteCategory 删除Category
//@Summary 删除Category
//@Description 删除Category
//@Tags category
//@Accept  json
//@Produce  json
//@Param id path string true "删除Category"
//@Success 200 {object} code.Response"{"code":200,"data":article.Category,"msg":"操作成功"}"
//@Router /category/delete [delete]
func (ct *CategoriesApi) DeleteCategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailWithMessage("id不能为空", c)
		return
	}
	if msg, err := categoryService.DeleteCategoryService(id); err != nil {
		code.FailWithMessage(msg, c)
		return
	} else {
		code.OkWithMessage(msg, c)
	}
}

//UpdateCategory 修改Category
//@Summary 修改Category
//@Description 修改Category
//@Tags category
//@Accept  json
//@Produce  json
//@Param category body article.Category true "创建Category"
//@Success 200 {object} code.Response"{"code":200,"data":article.Category,"msg":"操作成功"}"
//@Router /category/update [put]
func (ct *CategoriesApi) UpdateCategory(c *gin.Context) {
	var category article.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		code.FailWithMessage(err.Error(), c)
		return
	}
	if category.ID == "" {
		code.FailWithMessage("id不能为空", c)
		return
	}
	if msg, err := categoryService.UpdateCategoryService(category); err != nil {
		code.FailWithMessage(msg, c)
		return
	} else {
		code.OkWithMessage(msg, c)
	}

}

//GetCategories 查询Category
//@Summary 查询Category
//@Description 查询Category
//@Tags category
//@Accept  json
//@Produce  json
//@Success 200 {object} code.Response"{"code":200,"data":[]article.Category,"msg":"操作成功"}"
//@Router /category/list [get]
func (ct *CategoriesApi) GetCategories(c *gin.Context) {
	//keyword := c.Query("keyword")
	if list, msg, err := categoryService.GetCategoryService(); err != nil {
		code.FailWithMessage(msg, c)
		return
	} else {
		code.OkWithDetailed(list, msg, c)
	}
}
