package article

import (
	"gorm.io/gorm"
	"server/models"
	"server/models/system"
	"server/models/theme"
	"server/utils"
)

type Article struct {
	models.Model
	Title           string      `json:"title"  gorm:"type:varchar(100);not null" binding:"required"`            // 标题
	Status          string      `json:"status" gorm:"type:varchar(10);not null" binding:"required"`             // 状态 DRAFT, PUBLISHED
	Slug            string      `json:"slug" gorm:"type:varchar(100);"`                                         // 别名
	EditorType      string      `json:"editor_type" gorm:"type:varchar(100);"`                                  // 编辑器类型
	MetaKeyWords    string      `json:"meta_key_words" gorm:"type:varchar(100);"`                               // 头部关键字
	MetaDescription string      `json:"meta_description" gorm:"type:varchar(100);"`                             // 头部描述
	Summary         string      `json:"summary" gorm:"type:varchar(100);"`                                      // 摘要
	Thumbnail       string      `json:"thumbnail" gorm:"type:varchar(100);"`                                    // 缩略图
	DisableComments bool        `json:"disable_comments" gorm:"type:tinyint(1);"`                               // 禁止评论
	Password        string      `json:"password" gorm:"type:varchar(100);"`                                     // 访问密码
	WordCount       int         `json:"word_count" gorm:"type:int(10);"`                                        // 字数
	MdContent       string      `json:"md_content" gorm:"type:longblob;"`                                       // markdown内容
	HtmlContent     string      `json:"html_content" gorm:"type:longblob;"`                                     // html内容
	CommentCount    int         `json:"comment_count" gorm:"type:int(10);"`                                     // 评论数
	TagsID          []uint      `json:"tags_id" gorm:"-"`                                                       // tags id
	CategoriesID    []uint      `json:"categories_id" gorm:"-"`                                                 // 分类id
	IsTop           bool        `json:"is_top" gorm:"type:tinyint(1);"`                                         // 是否置顶
	Tags            []Tag       `gorm:"many2many:tag_articles;joinForeignKey:ArticleID" json:"tags"`            // tags
	Categories      []Category  `gorm:"many2many:category_articles;joinForeignKey:ArticleID" json:"categories"` // 分类
	Visits          int         `json:"visits" gorm:"type:int(10);"`                                            // 访问量
	Likes           int         `json:"likes" gorm:"type:int(10);"`                                             // 点赞数
	UserID          string      `json:"user_id" gorm:"type:varchar(100);comment:作者的UUID"`                       // 用户id
	User            system.User `json:"user"`                                                                   // 作者
	ThemeID         string      `json:"theme_id" gorm:"default:'default'"`                                      //主题ID
	Theme           theme.Theme `json:"theme"`                                                                  //主题
	HaloID          string      `json:"halo_id" gorm:"size:30;comment:halo的文章ID"`                               // halo的文章ID
}

type ListArticleRequest struct {
	Title      string `json:"title"` // 标题
	StartTime  string `form:"start_time" json:"start_time"`
	EndTime    string `form:"end_time" json:"end_time"`
	Status     string `form:"status" json:"status"`
	CategoryID int    `form:"category_id" json:"category_id"`
	TagID      int    `form:"tag_id" json:"tag_id"`
}

func init() {
	utils.AddAutoMigrateMethods(func(db *gorm.DB) {
		var tagCount int64
		if err := db.Model(&Tag{}).Count(&tagCount).Error; err != nil {
			return
		}
		if tagCount == 0 {
			tag := []Tag{
				{
					Name:      "默认标签",
					Slug:      "默认标签",
					Thumbnail: "",
					Color:     "red",
				},
			}
			if err := db.Create(&tag).Error; err != nil {
				return
			}
		}
		var categoryCount int64
		if err := db.Model(&Category{}).Count(&categoryCount).Error; err != nil {
			return
		}
		if categoryCount == 0 {
			category := []Category{
				{
					Name:        "默认分类",
					Slug:        "默认分类",
					Description: "默认分类描述",
					Thumbnail:   "",
				},
			}
			if err := db.Create(&category).Error; err != nil {
				return
			}
		}
	})
}
