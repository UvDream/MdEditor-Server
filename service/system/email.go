package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/global"
	"server/utils"
	"time"
)

func (*SysUserService) SendEmailCodeService(c *gin.Context, email string) (uniqueVerificationCode string, cd int, err error) {
	redis := global.Redis
	//生成验证码
	verificationCode := utils.GenerateVerificationCode()
	//生成唯一id
	uniqueVerificationCode = utils.GenerateUniqueVerificationCode()
	//先查询该邮箱是否已经发送过验证码
	_, err = redis.Get(c, email).Result()
	if err == nil {
		return "", code.EmailHasSend, err
	}
	//将邮箱存入redis,五分钟只能发送一次
	if err := redis.Set(c, email, verificationCode, 5*time.Minute).Err(); err != nil {
		return "", code.ErrorRedisSet, err
	}
	//验证码和唯一id存入redis,有效期五分钟
	if err := redis.Set(c, uniqueVerificationCode, verificationCode, 5*time.Minute).Err(); err != nil {
		return "", code.ErrorRedisSet, err
	}
	htmlBody := "您的验证码为：" + verificationCode + "，有效期为5分钟，请尽快使用。"
	//发送邮件
	if err := utils.SendEmail(email, htmlBody, email); err != nil {
		return "", code.ErrorSendEmail, err
	}
	return uniqueVerificationCode, code.SUCCESS, nil
}

func (*SysUserService) VerifyEmailCodeService(c *gin.Context, verificationCode string, uniqueVerificationCode string, email string) (cd int, err error) {
	redis := global.Redis
	//先查询该邮箱是否已经发送过验证码
	_, err = redis.Get(c, email).Result()
	if err != nil {
		return code.EmailHasNotSend, err
	}
	//从redis中获取验证码
	msg, err := redis.Get(c, uniqueVerificationCode).Result()
	if err != nil {
		return code.ErrorVerificationCode, err
	}
	//判断验证码是否正确
	if msg != verificationCode {
		return code.ErrorVerificationCode, err
	}
	return code.SUCCESS, nil
}
