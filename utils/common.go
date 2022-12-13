package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

// Paginator 分页查询处理
func Paginator(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	page := getParam(c, "page", 1)
	pageSize := getParam(c, "page_size", 10)

	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}

// RangeTime 时间查询处理
func RangeTime(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	startTime, _ := c.GetQuery("start_time")
	endTime, _ := c.GetQuery("end_time")
	return func(db *gorm.DB) *gorm.DB {
		if startTime != "" && endTime != "" {
			return db.Where("create_time between ? and ?", startTime, endTime)
		}
		return db
	}
}

func getParam(c *gin.Context, key string, defaultValue int) int {
	var result = defaultValue
	if param, exists := c.GetQuery(key); exists {
		var err error
		result, err = strconv.Atoi(param)
		if err != nil {
			result = defaultValue
		}
	}
	return result
}

// RandString 生成随机字符串
func RandString(n int) string {
	var letter = []byte("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result = make([]byte, n)
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}

// GenerateVerificationCode 生成随机验证码
func GenerateVerificationCode() string {
	var letter = []byte("123456789")
	var result = make([]byte, 6)
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}

func GenerateUniqueVerificationCode() string {
	return RandString(32)
}
