package code

var zhCNText = map[int]string{
	SUCCESS:                             "成功",
	ERROR:                               "错误",
	ErrorNotLogin:                       "未登录",
	ErrorTokenGenerate:                  "token生成失败",
	ErrorAuthorizationExpires:           "授权过期",
	ErrorImageNotFound:                  "图片不存在",
	UploadImageSuccess:                  "上传图片成功",
	ErrorCreateDir:                      "创建文件夹失败",
	SaveFileSuccess:                     "保存文件成功",
	ErrorSaveFile:                       "保存文件失败",
	ErrorSaveFileData:                   "保存文件信息到数据库失败",
	ErrorUploadQiNiu:                    "七牛云上传失败",
	ErrorOpenFile:                       "打开文件失败",
	ErrorUploadQiNiuSuccess:             "七牛云上传成功",
	ErrorMissingId:                      "缺少id",
	ErrorFileNotFound:                   "文件不存在",
	ErrorDeleteFile:                     "删除文件失败",
	ErrorDeleteFileData:                 "删除文件数据库失败",
	ErrorListFile:                       "获取文件列表失败",
	ErrorCreateTheme:                    "创建主题失败",
	ErrorGetTheme:                       "获取主题失败",
	ErrorDeleteTheme:                    "删除主题失败",
	ErrorThemeNotFound:                  "主题不存在",
	ErrorUpdateTheme:                    "主题更新失败",
	ErrorRegisterMissingParam:           "注册用户缺少参数",
	ErrorUserExist:                      "用户已存在",
	ErrorUserExistAccount:               "账号已存在",
	ErrorUserExistNickname:              "昵称已存在",
	ErrorUserExistEmail:                 "邮箱已存在",
	ErrorUserExistPhone:                 "手机号已存在",
	ErrorCreateUser:                     "创建用户失败",
	ErrorGetQueryParam:                  "获取查询参数失败",
	ErrorNotFound:                       "未找到",
	ErrorEditorTheme:                    "编辑主题失败",
	ErrorFindArticle:                    "查询文章失败",
	ErrLedgerExist:                      "账本已存在",
	ErrCreateLedger:                     "创建账本失败",
	ErrGetLedgerList:                    "获取账本列表失败",
	ErrCreateLedgerCategory:             "创建账本分类失败",
	ErrCreateLedgerCategoryRelation:     "创建账本分类关系失败",
	ErrorDeleteLedger:                   "删除账本失败",
	ErrorLedgerNotExist:                 "账本不存在",
	ErrorGetLedgerCategoryList:          "查询账本分类失败",
	ErrorGetLedgerCategoryRelationList:  "查询账本分类关联关系失败",
	ErrorDeleteLedgerCategoryRelation:   "删除账本分类关联关系失败",
	ErrorDeleteLedgerCategory:           "删除账本分类失败",
	ErrorUpdateLedger:                   "更新账本失败",
	ErrorGetLedger:                      "查询账本失败",
	HaloServerError:                     "halo请求错误",
	HaloTokenExpired:                    "halo登录过期",
	ErrorFindUser:                       "查询用户失败",
	ErrorUpYunPut:                       "又拍云上传失败",
	ErrorSetUserConfig:                  "设置用户配置失败",
	ErrorFindUserConfig:                 "查询用户配置失败",
	ErrorExportArticle:                  "文章导出错误",
	ErrorSkiUpload:                      "ski图床上传失败",
	ErrorReadThirdPartyResponse:         "读取第三方响应失败",
	ErrorSkiDelete:                      "ski图床删除失败",
	ErrorToken:                          "token错误",
	ErrorUpdatePasswordMissingParam:     "更新密码参数有误",
	ErrorCreateCategory:                 "创建分类失败",
	ErrCategoryExist:                    "分类已存在",
	ErrorUpdateCategory:                 "更新分类失败",
	ErrorCategoryNotExist:               "分类不存在",
	ErrorDeleteCategory:                 "删除分类失败",
	ErrCreateBill:                       "创建账单失败",
	ErrorDeleteBill:                     "删除账单失败",
	ErrorUpdateBill:                     "更新账单失败",
	ErrorGetBill:                        "查询账单失败",
	ErrorMissingLedgerId:                "缺少账本id",
	ErrCreateBudget:                     "创建预算失败",
	ErrDeleteBudget:                     "删除预算失败",
	ErrGetBudget:                        "查询预算失败",
	ErrUpdateBudget:                     "更新预算失败",
	ErrBudgetNotExist:                   "预算不存在",
	ErrBudgetExist:                      "预算已存在",
	ErrorGetCategory:                    "查询分类失败",
	ErrorGetCategoryStatistics:          "查询分类统计失败",
	ErrorGetIncomeExpenditureStatistics: "查询收支统计失败",
	ErrorNotLedgerCreator:               "不是账本创建者",
	ErrorShareLedger:                    "分享账本失败",
	ErrShareCodeExpired:                 "分享码已过期",
	ErrShareCodeNotExist:                "分享码不存在",
	ErrorJoinLedger:                     "加入账本失败",
	ErrorUserNotExist:                   "用户不存在",
	ErrorUserAlreadyJoined:              "用户已加入",
	ErrorGetBudget:                      "查询预算统计失败",
	ErrCreateLedgerNotMember:            "创建账本失败，不是账本成员",
	ErrorCreateLedgerNotMember:          "非会员只能创建一个账本",
	ErrorCreateLedgerMember:             "会员最大能创建五个账本",
	ErrorEmailMissingParam:              "邮箱参数有误",
	ErrorSendEmail:                      "发送邮件失败",
	EmailHasSend:                        "邮件已发送,五分钟内只能发送一次",
	ErrorRedisGet:                       "redis获取失败",
	ErrorVerificationCode:               "验证码错误",
	EmailHasNotSend:                     "邮件未发送",
	ErrorUpdateUser:                     "更新用户失败",
	ErrorEmailExist:                     "邮箱已存在",
	ErrorSetUserInviteCode:              "设置用户邀请码失败",
	ErrorGetCategoryStatisticsDetail:    "查询分类统计详情失败",
	ErrorGetTotalAmount:                 "查询总金额失败",
	ErrorFeedback:                       "反馈失败",
	ErrorFeedbackList:                   "查询反馈列表失败",
	ErrorMissingSearchParam:             "缺少搜索参数",
	ErrorUserList:                       "获取用户列表失败",
	ErrorRoleList:                       "获取角色列表失败",
	ErrorGetUser:                        "获取用户失败",
	ErrorGetPermission:                  "获取权限失败",
	ErrorGetLedgerMember:                "获取账本成员失败",
	ErrorGetIcon:                        "获取图标失败",
	ErrorGetColor:                       "获取颜色失败",
	ErrIcon:                             "图标错误",
	ErrIconClassification:               "新增icon分类失败",
	ErrorDeleteMenuFail:                 "删除菜单失败",
	ErrorAddMenuFail:                    "新增菜单失败",
	ErrorMenuList:                       "获取菜单列表失败",
	ErrorUserRoleFail:                   "分配用户角色失败",
	ErrorGetAppNeedUpdate:               "获取app版本失败",
}
