package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"server/global"
	"strconv"
)

func SendEmail(title string, body string, receiveEmail string) (err error) {
	sysConfig := global.Config.Email
	//接收邮箱设置
	mailTo := []string{
		receiveEmail,
	}
	fmt.Println(mailTo)
	//定义邮箱服务器信息
	mailConn := map[string]string{
		"user": sysConfig.Username,
		"pass": sysConfig.Password,
		"host": sysConfig.Host,
		"port": sysConfig.Port,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "[饭米粒记账]验证码")) //这种方式可以添加别名
	m.SetHeader("To", mailTo...)                                         //发送给多个用户
	m.SetHeader("Subject", title)                                        //设置邮件主题
	m.SetBody("text/html", body)                                         //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
