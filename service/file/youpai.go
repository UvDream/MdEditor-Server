package file

import (
	"github.com/upyun/go-sdk/v3/upyun"
	"go.uber.org/zap"
	"mime/multipart"
	code2 "server/code"
	"server/global"
	"server/models/system"
)

type YoupaiService struct{}

// UploadFile 上传文件到又拍云
func (*YoupaiService) UploadFile(fileHeader *multipart.FileHeader, file multipart.File, config system.UserConfig) (path string, key string, c int, err error) {
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   config.UpYunBucket,
		Operator: config.UpYunUser,
		Password: config.UpYunPass,
	})
	//上传文件
	if err := up.Put(&upyun.PutObjectConfig{
		Path:   fileHeader.Filename,
		Reader: file,
	}); err != nil {
		global.Log.Error("上传文件到又拍云失败", zap.Error(err))
		return "", "", code2.ErrorUpYunPut, err
	}
	//TODO 需要获取又拍云返回的文件地址
	return config.UpYunDomain + "/" + fileHeader.Filename, "", code2.SUCCESS, nil
}

// DeleteFile 删除又拍云文件
func (*YoupaiService) DeleteFile(key string, token string) error {
	return nil
}
