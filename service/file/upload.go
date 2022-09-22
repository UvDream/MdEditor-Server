package file

import (
	"mime/multipart"
	"server/models/system"
)

type OSS interface {
	UploadFile(fileHeader *multipart.FileHeader, file multipart.File, config system.UserConfig) (string, string, int, error)
	DeleteFile(key string) error
}

func NewOss(ossType string) OSS {
	switch ossType {
	case "local":
		return &LocalService{}
	case "qiniu":
		return &QiniuService{}
	case "youpai":
		return &YoupaiService{}
	default:
		return &LocalService{}
	}
}

func DeleteOss(position string) OSS {
	switch position {
	case "local":
		return &LocalService{}
	case "qiniu":
		return &QiniuService{}
	case "youpai":
		return &YoupaiService{}
	default:
		return &LocalService{}
	}
}
