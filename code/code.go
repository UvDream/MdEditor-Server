/*
 * @Author: wangzhongjie
 * @Date: 2022-08-12 15:21:50
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2023-09-25 17:19:24
 * @Description:
 * @Email: UvDream@163.com
 */
package code

import "server/global"

const (
	SUCCESS                         = 200
	ERROR                           = 500
	ErrorNotLogin                   = 50000
	ErrorAuthorizationExpires       = 50001
	ErrorImageNotFound              = 1001  //图片不存在
	UploadImageSuccess              = 2001  //上传图片成功
	ErrorCreateDir                  = 1002  //创建文件夹失败
	SaveFileSuccess                 = 2002  //保存文件成功
	ErrorSaveFile                   = 1003  //保存文件失败
	ErrorSaveFileData               = 1004  //保存文件信息到数据库失败
	ErrorUploadQiNiu                = 1005  //七牛云上传失败
	ErrorOpenFile                   = 1006  //打开文件失败
	ErrorUploadQiNiuSuccess         = 3001  //七牛云上传成功
	ErrorMissingId                  = 4001  //缺少id
	ErrorFileNotFound               = 404   //文件不存在
	ErrorDeleteFile                 = 5001  //删除文件失败
	ErrorDeleteFileData             = 5002  //删除文件数据库失败
	ErrorListFile                   = 6001  //获取文件列表失败
	ErrorCreateTheme                = 7001  //创建主题失败
	ErrorGetTheme                   = 8001  //获取主题失败
	ErrorDeleteTheme                = 9001  //删除主题失败
	ErrorThemeNotFound              = 10001 //主题不存在
	ErrorUpdateTheme                = 10002 //主题更新失败
	ErrorRegisterMissingParam       = 10003 //缺少参数
	ErrorUserExist                  = 10011 //用户已存在
	ErrorUserExistAccount           = 10012 //账户已存在
	ErrorUserExistNickname          = 10013 //昵称已存在
	ErrorUserExistEmail             = 10014 //邮箱已存在
	ErrorUserExistPhone             = 10015 //手机号已存在
	ErrorCreateUser                 = 10016 //用户创建失败
	ErrorGetQueryParam              = 10017 //获取查询参数失败
	ErrorNotFound                   = 10018 //未找到记录
	ErrorEditorTheme                = 10019 //编辑主题失败
	ErrorFindArticle                = 10020 //查询文章失败
	ErrorFindUser                   = 10021 //	查询用户信息失败
	ErrorUpYunPut                   = 10022 //又拍云上传失败
	ErrorSetUserConfigMissingParam  = 10023 //缺少参数
	ErrorSetUserConfig              = 10024 //设置用户配置失败
	ErrorFindUserConfig             = 10025 //查询用户配置失败
	ErrorExportArticle              = 10026 //导出文章失败
	ErrorSkiUpload                  = 10027 //上传图片失败
	ErrorReadThirdPartyResponse     = 10028 //读取第三方响应失败
	ErrorSkiDelete                  = 10029 //删除图片失败
	ErrorToken                      = 10030 //token错误
	ErrorUpdatePasswordMissingParam = 10031 //缺少参数
	ErrorUpdatePassword             = 10032 //更新密码失败
	ErrorTokenGenerate              = 10033 //token生成错误
	ErrorGetOpenID                  = 10034 //获取openID 失败
	ErrorEmailMissingParam          = 10035 //邮件验证码缺少参数
	ErrorRedisSet                   = 10036 //redis set失败
	ErrorSendEmail                  = 10037 //发送邮件失败
	EmailHasSend                    = 10038 //邮件已发送
	ErrorRedisGet                   = 10039 //redis get失败
	ErrorVerificationCode           = 10040 //验证码错误
	EmailHasNotSend                 = 10041 //邮件未发送
	ErrorUserNotExist               = 10042 //用户不存在
	ErrorUpdateUser                 = 10043 //更新用户失败
	ErrorEmailExist                 = 10044 //邮箱已存在
	ErrorSetUserInviteCode          = 10045 //设置用户邀请码失败
	ErrorFeedback                   = 10046 //反馈保存失败
	ErrorFeedbackList               = 10047 //反馈列表获取失败
	ErrorUserList                   = 10048 // 获取用户列表失败
	ErrorRoleList                   = 10049 //获取角色列表失败
	ErrorGetUser                    = 10050 //获取用户失败
	ErrIconClassification           = 10051 //新增icon分类失败
	ErrorDeleteMenuFail             = 10052 //删除菜单失败
	ErrorAddMenuFail                = 10053 //新增菜单失败
	ErrorMenuList                   = 10054 //菜单列表获取失败
	ErrorUserRoleFail               = 10055 //分配用户角色失败
	ErrorGetAppNeedUpdate           = 10056 //获取app版本失败
	ErrorLoopAccount                = 10057 //周期记账失败
	ErrorDeleteLoopAccount          = 10058 //删除周期记账失败
	ErrorUpdateLoopAccount          = 10059 //更新周期记账失败
	ErrorGetLoopList                = 10060 //获取周期记账列表失败
	ErrorLoopAccountNotExist        = 10061 //周期记账不存在
)

func Text(code int) string {
	lang := global.Config.System.Language
	if lang == "en-us" {
		return enUSText[code]
	}
	if lang == "zh-cn" {
		return zhCNText[code]
	}
	return zhCNText[code]
}
