package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	ledger2 "server/models/ledger"
	"server/service/ledger"
	"server/utils"
)

func (*LedgerAdminService) GetColorListService(ledgerId string, isBgColor string, c *gin.Context) (colorList []ledger2.Color, total int64, cd int, err error) {
	userList, cd, err := ledger.GetLedgerUserListService(ledgerId)
	if err != nil {
		return colorList, total, cd, err
	}
	db := global.DB
	if isBgColor != "" {
		db = db.Where("is_bg_color = ?", isBgColor)
	}

	if err := db.Where("user_id in ?", userList).Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&colorList).Error; err != nil {
		return colorList, total, cd, err
	}
	return colorList, total, cd, err
}

func (*LedgerAdminService) GetIconListService(ledgerId string, c *gin.Context) (IconList []ledger2.IconClassification, total int64, cd int, err error) {
	userList, cd, err := ledger.GetLedgerUserListService(ledgerId)
	if err != nil {
		return IconList, total, cd, err
	}
	db := global.DB
	if err := db.Where("user_id in ?", userList).Preload(clause.Associations).Scopes(utils.Paginator(c)).Find(&IconList).Error; err != nil {
		return IconList, total, cd, err
	}
	return IconList, total, cd, err
}

func (*LedgerAdminService) AddColorService(color ledger2.Color, c *gin.Context) (cd int, err error) {
	userID := utils.FindUserID(c)
	color.UserID = userID
	if err := global.DB.Create(&color).Error; err != nil {
		return code.ErrColor, err
	}
	return code.SUCCESS, err
}