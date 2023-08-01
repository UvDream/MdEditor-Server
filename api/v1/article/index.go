package article

import "server/service"

type ApiArticleGroup struct {
	ArticlesApi
	CategoriesApi
	TagsApi
}

var (
	articleService  = service.ServicesGroupApp.ToArticleService
	tagService      = service.ServicesGroupApp.TagService
	categoryService = service.ServicesGroupApp.CategoryService
)
