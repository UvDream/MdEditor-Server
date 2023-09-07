package admin

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/ledger"
	"server/utils"
)

// GetColorList 获取颜色列表
// @Summary 获取颜色列表
// @Tags admin/color
// @Accept  json
// @Produce  json
// @Param ledger_id query int true "账单ID"
// @Param is_bg_color query string false "是否是背景颜色"
// @Param query query models.PaginationRequest true "分页参数"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool}
// @Router /admin/ledger/color/list [get]
func (*LedgerAdminApi) GetColorList(c *gin.Context) {
	ledgerId := c.Query("ledger_id")
	isBgColor := c.Query("is_bg_color")
	data, total, cd, err := ledgerAdminService.GetColorListService(ledgerId, isBgColor, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}

// AddColor 新增颜色
// @Summary 新增颜色
// @Tags admin/color
// @Accept  json
// @Produce  json
// @Param color body ledger.Color true "颜色"
// @Success 200 {object} code.Response
// @Router /admin/ledger/color/add [post]
func (*LedgerAdminApi) AddColor(c *gin.Context) {
	var color ledger.Color
	if err := c.ShouldBindJSON(&color); err != nil {
		code.FailResponse(code.ErrColor, c)
		return
	}
	cd, err := ledgerAdminService.AddColorService(color, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// GetIconList 获取图标列表admin
// @Summary 获取图标列表admin
// @Tags admin/icon
// @Accept  json
// @Produce  json
// @Param ledger_id query int true "账单ID"
// @Param query query models.PaginationRequest true "分页参数"
// @Success 200 {object} code.Response
// @Router /admin/ledger/icon/list [get]
func (*LedgerAdminApi) GetIconList(c *gin.Context) {
	ledgerId := c.Query("ledger_id")
	data, total, cd, err := ledgerAdminService.GetIconListService(ledgerId, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}

// AddIconClassification 新增icon分类
// @Summary 新增icon分类
// @Tags admin/icon
// @Accept  json
// @Produce  json
// @Param icon_classification body ledger.IconClassification true "icon分类"
// @Success 200 {object} code.Response{}
// @Router /admin/ledger/icon/classification/add [post]
func (*LedgerAdminApi) AddIconClassification(c *gin.Context) {
	var iconClassification ledger.IconClassification
	if err := c.ShouldBindJSON(&iconClassification); err != nil {
		code.FailResponse(code.ErrIconClassification, c)
		return
	}
	userId := utils.FindUserID(c)
	iconClassification.UserID = userId
	cd, err := ledgerAdminService.AddIconClassificationService(iconClassification)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// AddIcon 新增icon
// @Summary 增加图标
// @Tags admin/icon
// @Accept json
// @Produce json
// @Param icon body []ledger.Icon true "图标"
// @Success 200 {object} code.Response{}
// @Router /admin/ledger/icon/add [post]
func (*LedgerAdminApi) AddIcon(c *gin.Context) {
	var icons []ledger.Icon
	if err := c.ShouldBindJSON(&icons); err != nil {
		code.FailResponse(code.ErrIcon, c)
		return
	}
	userId := utils.FindUserID(c)
	for i := 0; i < len(icons); i++ {
		icons[i].UserID = userId
	}
	cd, err := ledgerAdminService.AddIconService(icons)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// DeleteIconClassification 删除icon分类
// @Summary 删除icon分类
// @Tags admin/icon
// @Accept  json
// @Produce  json
// @Param id query string true "icon分类ID"
// @Success 200 {object} code.Response
// @Router /admin/ledger/icon/classification/delete [delete]
func (*LedgerAdminApi) DeleteIconClassification(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerAdminService.DeleteIconClassificationService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// DeleteIcon 删除图标
// @Summary 删除图标
// @Tags admin/icon
// @Accept  json
// @Produce  json
// @Param id query string true "图标ID"
// @Success 200 {object} code.Response
// @Router /admin/ledger/icon/delete [delete]
func (*LedgerAdminApi) DeleteIcon(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerAdminService.DeleteIconService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}

// DeleteColor 删除颜色
// @Summary 删除颜色
// @Tags admin/color
// @Accept  json
// @Produce  json
// @Param id query string true "颜色ID"
// @Success 200 {object} code.Response
// @Router /admin/ledger/color/delete [delete]
func (*LedgerAdminApi) DeleteColor(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerAdminService.DeleteColorService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}
