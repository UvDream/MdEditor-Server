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

type LedgerAdminService struct{}

func (*LedgerAdminService) GetLedgerListService(query ledger.LedgerRequest, c *gin.Context) (data []ledger.Ledger, total int64, cd int, err error) {
	userId := utils.FindUserID(c)
	var userService service.SysUserService
	roles, err := userService.FindUserRoles(userId)
	if err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	if len(roles) > 0 {
		for _, v := range roles {
			if v == "admin" || v == "root" {
				data, total, cd, err = getAllLedgerList(query, c)
				return data, total, cd, err
			}
		}
	}
	//普通用户
	data, total, cd, err = getNormalLedgerList(query, c)
	if err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	return data, total, code.SUCCESS, err
}

func getNormalLedgerList(query ledger.LedgerRequest, c *gin.Context) (data []ledger.Ledger, total int64, cd int, err error) {
	userId := utils.FindUserID(c)
	db := global.DB
	dp := global.DB
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}
	if query.Description != "" {
		db = db.Where("description LIKE ?", "%"+query.Description+"%")
	}
	//开始时间和结束时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	//查询自己创建的账本
	if err := db.Where("creator_id = ?", userId).Preload(clause.Associations).Find(&data).Error; err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	//查询用户协同的账本
	var ledgerUser []ledger.LedgerUser
	if err := dp.Where("user_id = ?", userId).Find(&ledgerUser).Error; err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	//查询协同账本
	for _, v := range ledgerUser {
		var ledger ledger.Ledger
		if err := db.Where("id = ?", v.LedgerID).Preload(clause.Associations).Find(&ledger).Error; err != nil {
			return nil, 0, code.ErrorMissingLedgerId, err
		}
		data = append(data, ledger)
	}
	total = int64(len(data))
	return data, total, code.SUCCESS, err
}

func getAllLedgerList(query ledger.LedgerRequest, c *gin.Context) (data []ledger.Ledger, total int64, cd int, err error) {
	db := global.DB
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}
	if query.Description != "" {
		db = db.Where("description LIKE ?", "%"+query.Description+"%")
	}
	//开始时间和结束时间
	if query.StartTime != "" && query.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", query.StartTime, query.EndTime)
	}
	//查询总数
	if err := db.Model(&ledger.Ledger{}).Count(&total).Error; err != nil {
		return nil, 0, code.ErrorMissingLedgerId, err
	}
	//查询数据
	if err := db.Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&data).Error; err != nil {
		return nil, 0, code.ErrorLedgerNotExist, err
	}
	return data, total, code.SUCCESS, err
}
