package article

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"server/code"
	"server/global"
	article2 "server/models/article"
	"server/utils"
)

func (*ToArticleService) ExportMd(userID string, articleID string) (filePath string, cd int, err error) {
	//	首先查出文章
	db := global.DB
	var article article2.Article
	if err := db.Where("id = ?", articleID).Find(&article).Error; err != nil {
		return "", code.ErrorFindArticle, err
	}
	//查询文章是否是你的
	if article.UserID != userID {
		return "", code.ErrorExportArticle, err
	}
	//导出文章
	//写入的文章目录
	dir := global.Config.Local.Path + "/markdown/"
	//首先判断文件夹是否存在
	if err := fileutil.IsExist(dir); !err {
		if err := fileutil.CreateDir(dir); err != nil {
			return "", code.ErrorMissingId, err
		}
	}
	path, err := MkMarkdown(dir, article)
	if err != nil {
		return "", code.ErrorExportArticle, err
	}
	return path, code.SUCCESS, nil
}

func MkMarkdown(path string, article article2.Article) (string, error) {
	//	创建文件
	fileName := article.Title + "_" + article.ID + ".md"
	filePath := path + fileName
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
		return filePath, err
	}
	defer file.Close()
	//	写入文件
	_, err = file.WriteString(article.MdContent)
	if err != nil {
		return filePath, err
	}
	return utils.SimplifyPath(filePath), nil
}
