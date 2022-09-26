package system

import "server/models"

type UserConfig struct {
	models.Model
	//存储位置
	OSSType string `json:"oss_type" gorm:"comment:存储位置"`
	//是否是https
	IsHttps bool `json:"is_https" gorm:"comment:是否是https"`
	//	七牛云
	QiNiuAccessKey string `json:"qi_niu_access_key" gorm:"comment:七牛云AccessKey"`
	QiNiuSecretKey string `json:"qi_niu_secret_key" gorm:"comment:七牛云SecretKey"`
	QiNiuBucket    string `json:"qi_niu_bucket" gorm:"comment:七牛云Bucket"`
	QiNiuDomain    string `json:"qi_niu_domain" gorm:"comment:七牛云Domain"`
	QiNiuPosition  string `json:"qi_niu_position" gorm:"comment:七牛云存储位置"`
	//	又拍云
	UpYunBucket string `json:"up_yun_bucket" gorm:"comment:又拍云Bucket"`
	UpYunDomain string `json:"up_yun_domain" gorm:"comment:又拍云Domain"`
	UpYunUser   string `json:"up_yun_user" gorm:"comment:又拍云User"`
	UpYunPass   string `json:"up_yun_pass" gorm:"comment:又拍云Pass"`
}
