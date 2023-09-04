/*
 * @Author: wangzhongjie
 * @Date: 2022-08-12 15:21:50
 * @LastEditors: wangzhongjie
 * @LastEditTime: 2023-07-28 17:25:45
 * @Description:
 * @Email: UvDream@163.com
 */
package code

var enUSText = map[int]string{
	SUCCESS:                            "Success",
	ERROR:                              "error",
	ErrorNotLogin:                      "not login",
	ErrorAuthorizationExpires:          "Authorization Expires",
	ErrorImageNotFound:                 "Image not found",
	UploadImageSuccess:                 "Upload image success",
	ErrorCreateDir:                     "Create dir failed",
	SaveFileSuccess:                    "Save file success",
	ErrorSaveFile:                      "Save file failed",
	ErrorSaveFileData:                  "Save file data failed",
	ErrorUploadQiNiu:                   "Upload to qiniu failed",
	ErrorOpenFile:                      "Open file failed",
	ErrorUploadQiNiuSuccess:            "Upload to qiniu success",
	ErrorMissingId:                     "Missing id",
	ErrorFileNotFound:                  "File not found",
	ErrorDeleteFile:                    "Delete file failed",
	ErrorDeleteFileData:                "Delete file data failed",
	ErrorListFile:                      "List file failed",
	ErrorCreateTheme:                   "Create theme failed",
	ErrorGetTheme:                      "Get theme failed",
	ErrorDeleteTheme:                   "Delete theme failed",
	ErrorThemeNotFound:                 "Theme not found",
	ErrorUpdateTheme:                   "Update theme failed",
	ErrorRegisterMissingParam:          "Missing param",
	ErrorUserExist:                     "User exist",
	ErrorUserExistAccount:              "User exist ledger",
	ErrorUserExistNickname:             "User exist nickname",
	ErrorUserExistEmail:                "User exist email",
	ErrorUserExistPhone:                "User exist phone",
	ErrorCreateUser:                    "Create user failed",
	ErrorGetQueryParam:                 "Get query param failed",
	ErrorNotFound:                      "Not found",
	ErrorEditorTheme:                   "Editor theme failed",
	ErrorFindArticle:                   "find article failed",
	ErrLedgerExist:                     "ledger exist",
	ErrCreateLedger:                    "create ledger failed",
	ErrGetLedgerList:                   "get ledger list failed",
	ErrCreateLedgerCategory:            "create ledger category failed",
	ErrCreateLedgerCategoryRelation:    "create ledger category relation failed",
	ErrorDeleteLedger:                  "delete ledger failed",
	ErrorLedgerNotExist:                "ledger not exist",
	ErrorGetLedgerCategoryList:         "get ledger category list failed",
	ErrorGetLedgerCategoryRelationList: "get ledger category relation list failed",
	ErrorDeleteLedgerCategoryRelation:  "delete ledger category relation failed",
	ErrorDeleteLedgerCategory:          "delete ledger category failed",
	ErrorUpdateLedger:                  "update ledger failed",
	ErrorGetLedger:                     "get ledger failed",
	HaloServerError:                    "Halo server error",
	HaloTokenExpired:                   "Halo token expired",
	ErrorFindUser:                      "find user failed",
	ErrorUpYunPut:                      "upyun put failed",
	ErrorSetUserConfig:                 "set user config failed",
	ErrorFindUserConfig:                "find user config failed",
	ErrorExportArticle:                 "export article failed",
	ErrorSkiUpload:                     "ski upload failed",
	ErrorReadThirdPartyResponse:        "read third party response failed",
	ErrorSkiDelete:                     "ski delete failed",
	ErrorToken:                         "token error",
	ErrorUpdatePasswordMissingParam:    "update password missing param",
	ErrorCreateCategory:                "create category failed",
	ErrCategoryExist:                   "category exist",
	ErrorUpdateCategory:                "update category failed",
	ErrorCategoryNotExist:              "category not exist",
	ErrorDeleteCategory:                "delete category failed",
	ErrorGetCategoryStatisticsDetail:   "get category statistics detail failed",
	ErrorGetTotalAmount:                "get total amount failed",
	ErrorFeedback:                      "feedback failed",
	ErrorFeedbackList:                  "feedback list failed",
	ErrorMissingSearchParam:            "missing search param",
	ErrorUserList:                      "user list failed",
	ErrorRoleList:                      "role list failed",
	ErrorGetUser:                       "get user failed",
	ErrorGetPermission:                 "get permission failed",
	ErrorGetLedgerMember:               "get ledger member failed",
	ErrorGetIcon:                       "get icon failed",
	ErrorGetColor:                      "get color failed",
	ErrIcon:                            "icon error",
	ErrIconClassification:              "icon classification error",
}
