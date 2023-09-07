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

func (*LedgerAdminService) AddIconService(icons []ledger2.Icon) (cd int, err error) {
	db := global.DB
	for _, v := range icons {
		//	先查询是否存在,不存在则创建
		var icon ledger2.Icon
		if err := db.Where("icon_name = ?", v.ClassName).Where("user_id = ?", v.UserID).First(&icon).Error; err != nil {
			if err := db.Create(&v).Error; err != nil {
				return code.ErrIcon, err
			}
		}

	}
	return code.SUCCESS, err
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
	return IconList, total, code.SUCCESS, err
}

func (*LedgerAdminService) AddIconClassificationService(iconClassification ledger2.IconClassification) (cd int, err error) {
	db := global.DB
	if err := db.Create(&iconClassification).Error; err != nil {
		return code.ErrIcon, err
	}
	return code.SUCCESS, err
}

func (*LedgerAdminService) AddColorService(color ledger2.Color, c *gin.Context) (cd int, err error) {
	userID := utils.FindUserID(c)
	color.UserID = userID
	if err := global.DB.Create(&color).Error; err != nil {
		return code.ErrColor, err
	}
	return code.SUCCESS, err
}

// DeleteColorService 删除颜色
func (*LedgerAdminService) DeleteColorService(id string) (cd int, err error) {
	if err := global.DB.Where("id = ?", id).Delete(&ledger2.Color{}).Error; err != nil {
		return code.ErrColor, err
	}
	return code.SUCCESS, err
}

// DeleteIconClassificationService 删除icon分类
func (*LedgerAdminService) DeleteIconClassificationService(id string) (cd int, err error) {
	if err := global.DB.Where("id = ?", id).Delete(&ledger2.IconClassification{}).Error; err != nil {
		return code.ErrIconClassification, err
	}
	//删除所有分类下的icon
	if err := global.DB.Where("icon_classification_id = ?", id).Delete(&ledger2.Icon{}).Error; err != nil {
		return code.ErrIcon, err
	}
	return code.SUCCESS, err
}

// DeleteIconService 删除icon
func (*LedgerAdminService) DeleteIconService(id string) (cd int, err error) {
	if err := global.DB.Where("id = ?", id).Delete(&ledger2.Icon{}).Error; err != nil {
		return code.ErrIcon, err
	}
	return code.SUCCESS, err
}
