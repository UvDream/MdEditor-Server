package file

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"server/code"
	"server/global"
	file2 "server/models/file"
	"server/models/system"
	"server/utils"
	"strings"
)

type FilesService struct{}

// UploadFileService 文件上传
func (*FilesService) UploadFileService(c *gin.Context) (data file2.File, ce int, err error) {
	file, fileHeader, _ := c.Request.FormFile("file")
	userID := utils.FindUserID(c)
	//查询用户配置
	var user system.User
	if err := global.DB.Preload("UserConfig").Where("id = ?", userID).First(&user).Error; err != nil {
		return data, code.ErrorFindUser, err
	}
	token := c.Query("token")
	user.UserConfig.Token = token
	url, key, codes, err := NewOss(user.UserConfig.OSSType).UploadFile(fileHeader, file, user.UserConfig)
	if err != nil {
		return data, codes, err
	}
	data, ce, err = SaveFileData(fileHeader, url, key, userID, user.UserConfig.OSSType, user.UserConfig.IsHttps)
	return data, ce, err
}
func (*FilesService) PublicUploadFileService(c *gin.Context) (data file2.File, ce int, err error) {
	file, fileHeader, _ := c.Request.FormFile("file")
	userID := utils.FindUserID(c)
	var user system.User
	token := c.Query("token")
	user.UserConfig.Token = token
	user.UserConfig.OSSType = global.Config.System.OssType
	//qi_niu_access_key
	user.UserConfig.QiNiuAccessKey = global.Config.Qiniu.AccessKey
	//qi_niu_secret_key
	user.UserConfig.QiNiuSecretKey = global.Config.Qiniu.SecretKey
	//qi_niu_bucket
	user.UserConfig.QiNiuBucket = global.Config.Qiniu.Bucket
	//qi_niu_domain
	user.UserConfig.QiNiuDomain = global.Config.Qiniu.DomainName
	//qi_niu_position
	user.UserConfig.QiNiuPosition = global.Config.Qiniu.Domain
	user.UserConfig.IsHttps = global.Config.Qiniu.DomainProtocol == "https"
	url, key, codes, err := NewOss(user.UserConfig.OSSType).UploadFile(fileHeader, file, user.UserConfig)
	if err != nil {
		return data, codes, err
	}
	data, ce, err = SaveFileData(fileHeader, url, key, userID, user.UserConfig.OSSType, user.UserConfig.IsHttps)
	return data, ce, err
}

func (*FilesService) DeleteFileService(c *gin.Context, id string) (file file2.File, codeNumber int, err error) {
	db := global.DB
	if err := db.Where("id = ?", id).First(&file).Error; err != nil {
		return file, code.ErrorFileNotFound, err
	}
	// 删除os 文件
	token := c.Query("token")
	err = DeleteOss(file.Position).DeleteFile(file.Key, token)
	if err != nil {
		return file, code.ErrorDeleteFile, err
	}
	//删除数据库文件
	if err := db.Where("id = ?", id).Delete(&file).Error; err != nil {
		return file, code.ErrorDeleteFileData, err
	}
	return file, code.SUCCESS, err
}

// SaveFileData 保存数据到数据库
func SaveFileData(fileHeader *multipart.FileHeader, url string, key string, authID string, position string, IsHttps bool) (data file2.File, ce int, err error) {
	db := global.DB
	var file file2.File
	ext := path.Ext(fileHeader.Filename)
	file.Name = strings.TrimSuffix(fileHeader.Filename, ext)
	file.URL = url
	file.Size = fileHeader.Size
	file.Type = fileHeader.Header.Get("Content-Type")
	file.Position = position
	file.Key = key
	file.AuthID = authID
	file.IsHttps = IsHttps
	if err := db.Create(&file).Error; err != nil {
		return file, code.ErrorSaveFileData, err
	}
	return file, code.SUCCESS, nil
}
