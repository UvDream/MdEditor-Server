package ledger

import (
	"gorm.io/gorm/clause"
	"server/code"
	"server/global"
	ledger2 "server/models/ledger"
)

func (*LedgersService) GetIconListService(ledgerId string) (IconList []ledger2.IconClassification, cd int, err error) {
	userList, cd, err := GetLedgerUserListService(ledgerId)
	if err != nil {
		return IconList, cd, err
	}
	db := global.DB
	if err := db.Where("user_id in ?", userList).Preload(clause.Associations).Find(&IconList).Error; err != nil {
		return IconList, cd, err
	}
	return IconList, code.SUCCESS, nil
}

// GetIconColorListService 获取颜色列表
func (*LedgersService) GetIconColorListService(ledgerId string, isBgColor string) (colorList []ledger2.Color, cd int, err error) {
	userList, cd, err := GetLedgerUserListService(ledgerId)
	if err != nil {
		return colorList, cd, err
	}
	db := global.DB
	if err := db.Where("user_id in ?", userList).Where("is_bg_color = ?", isBgColor).Find(&colorList).Error; err != nil {
		return colorList, cd, err
	}
	return colorList, code.SUCCESS, nil
}

func GetColorListByUserIdService(userId string, isBgColor string) (colorList []ledger2.Color, cd int, err error) {
	if err := global.DB.Where("user_id = ?", userId).Where("is_bg_color = ?", isBgColor).Find(&colorList).Error; err != nil {
		return colorList, code.ErrorGetColor, err
	}
	return colorList, code.SUCCESS, nil
}

// GetLedgerUserListService 根据ledger_id 查询账本的用户
func GetLedgerUserListService(ledgerId string) (userList []string, cd int, err error) {
	var userLedger []ledger2.LedgerUser
	if err := global.DB.Where("ledger_id = ?", ledgerId).Find(&userLedger).Error; err != nil {
		return userList, code.ErrorGetLedgerMember, err
	}
	for _, v := range userLedger {
		userList = append(userList, v.UserID)
	}
	//查询账本创建者
	var ledger ledger2.Ledger
	if err := global.DB.Where("id = ?", ledgerId).First(&ledger).Error; err != nil {
		return userList, code.ErrorGetLedger, err
	}
	userList = append(userList, ledger.CreatorID)
	return userList, code.SUCCESS, nil
}

// GetIconListByUserIdService 根据用户id查询Icon
func GetIconListByUserIdService(userId string) (iconList []ledger2.IconClassification, cd int, err error) {
	if err := global.DB.Where("user_id = ?", userId).Preload(clause.Associations).Find(&iconList).Error; err != nil {
		return iconList, code.ErrorGetIcon, err
	}
	return iconList, code.SUCCESS, nil
}
