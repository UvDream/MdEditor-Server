package utils

import (
	"gorm.io/gorm"
	"server/global"
)

// 自动初始化数据

var AutoMigrateMethods []func(db *gorm.DB)

func AddAutoMigrateMethods(method func(client *gorm.DB)) {
	AutoMigrateMethods = append(AutoMigrateMethods, method)
}
func AutoMigrate() {
	for _, method := range AutoMigrateMethods {
		method(global.DB)
	}
}
