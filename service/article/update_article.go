package article

import (
	"server/global"
	"server/models/article"
)

func (*ToArticleService) UpdateArticleService(articleOpts article.Article) (ac article.Article, msg string, err error) {
	//修改文章
	//	存储数据库
	if err := global.DB.Model(&ac).Where("id = ?", articleOpts.ID).Omit("CreateTime").Omit("UpdateTime").Updates(&articleOpts).Error; err != nil {
		return articleOpts, "修改文章失败", err
	}
	//存储redis
	msg, err = SetArticleRedis(articleOpts)
	if err != nil {
		return articleOpts, "修改文章存储redis失败", err
	}
	return articleOpts, "修改文章成功", nil
}
