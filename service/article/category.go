package article

import (
	"server/global"
	"server/models/article"
)

type CategoryService struct{}

//CreateCategoryService 增加category
func (*CategoryService) CreateCategoryService(category article.Category) (cat *article.Category, msg string, err error) {
	db := global.DB
	err = db.Where("name = ?", category.Name).First(&article.Category{}).Error
	if err == nil {
		return &category, "category已经存在", err
	}
	if err := db.Create(&category).Error; err != nil {
		return &category, "创建category失败", err
	}
	return &category, "创建category成功", err
}

//DeleteCategoryService 删除category
func (*CategoryService) DeleteCategoryService(id string) (msg string, err error) {
	db := global.DB
	var category article.Category
	//查询是否存在
	err = db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return "category不存在", err
	}
	//查询是否有子级
	err = db.Where("parent_id = ?", id).First(&article.Category{}).Error
	if err == nil {
		return "category有子级，不能删除", err
	}
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return "删除category失败", err
	}
	return "删除category成功", err
}

//UpdateCategoryService 更新category
func (*CategoryService) UpdateCategoryService(category article.Category) (msg string, err error) {
	db := global.DB
	err = db.Where("id = ?", category.ID).First(&article.Category{}).Error
	if err != nil {
		return "category不存在", err
	}
	if err := db.Where("id = ?", category.ID).Save(&category).Error; err != nil {
		return "更新category失败", err
	}
	return "更新成功", nil
}

//GetCategoryService 获取category
func (*CategoryService) GetCategoryService() (categoryList []article.Category, msg string, err error) {
	db := global.DB
	//if keyword != "" {
	//	db = db.Where("name LIKE ?", "%"+keyword+"%").Or("slug LIKE ?", "%"+keyword+"%")
	//}
	//先查询顶级
	err = db.Where("parent_id = ?", "").Find(&categoryList).Error
	if err != nil {
		return categoryList, "获取category顶级失败", err
	}
	for k, i := range categoryList {
		//查询子级
		categoryList[k].Children = findChildrenCategory(i)
	}
	return categoryList, "查询category列表成功", err
}
func findChildrenCategory(category article.Category) (children []article.Category) {
	db := global.DB
	if err := db.Model(&article.Category{}).Where("parent_id = ?", category.ID).Find(&children).Error; err != nil {
		return nil
	}
	for k, i := range children {
		children[k].Children = findChildrenCategory(i)
	}
	return children
}
