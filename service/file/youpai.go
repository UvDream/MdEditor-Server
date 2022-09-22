package file

import (
	"fmt"
	"github.com/upyun/go-sdk/v3/upyun"
	"mime/multipart"
	code2 "server/code"
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
	fmt.Println(up)
	//上传文件
	if err := up.Put(&upyun.PutObjectConfig{
		Path:   fileHeader.Filename,
		Reader: file,
	}); err != nil {
		return "", "", code2.ErrorUpYunPut, err
	}
	return config.UpYunDomain + "/" + fileHeader.Filename, "", code2.SUCCESS, nil
}

// DeleteFile 删除又拍云文件
func (*YoupaiService) DeleteFile(key string) error {
	return nil
}
