package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	"server/models/ledger"
	service "server/service/system"
	"server/utils"
)

func (*LedgerAdminService) GetLoopListService(c *gin.Context) (data []ledger.LoopAccount, total int64, cd int, err error) {
	userID := utils.FindUserID(c)
	var userService service.SysUserService
	roles, err := userService.FindUserRoles(userID)
	if err != nil {
		code.FailResponse(code.ErrorGetPermission, c)
		return
	}
	isAdmin := false
	//判断是否是管理员
	for _, role := range roles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}
	db := global.DB

	if isAdmin {
		//	查询所有
		//查询total
		if err := db.Model(&ledger.LoopAccount{}).Count(&total).Error; err != nil {
			return data, total, code.ErrorGetLoopList, err
		}
		if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&data).Error; err != nil {
			return data, total, code.ErrorGetLoopList, err
		}
	} else {
		//	查询当前用户账本
		var ledgerService LedgerAdminService
		userID := utils.FindUserID(c)
		ledgerIdArr, err := ledgerService.GetUserLedgerService(userID)
		if err != nil {
			return data, total, code.ErrorGetLoopList, err
		}
		//	查询total
		if err := db.Model(&ledger.LoopAccount{}).Where("ledger_id in ?", ledgerIdArr).Count(&total).Error; err != nil {
			return data, total, code.ErrorGetLoopList, err
		}
		if err := db.Where("ledger_id in ?", ledgerIdArr).Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&data).Error; err != nil {
			return data, total, code.ErrorGetLoopList, err
		}
	}
	return data, total, code.SUCCESS, err
}
