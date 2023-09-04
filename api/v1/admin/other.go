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

// AddIcon 新增icon
// @Summary 增加图标
// @Tags admin/icon
// @Accept json
// @Produce json
// @Success 200 {object} code.Response
// @Router /admin/ledger/icon/add [post]
func (*LedgerAdminApi) AddIcon(c *gin.Context) {
	var icon ledger.Icon
	if err := c.ShouldBindJSON(&icon); err != nil {
		code.FailResponse(code.ErrIcon, c)
		return
	}
	userId := utils.FindUserID(c)
	icon.UserID = userId
	cd, err := ledgerAdminService.AddIconService(icon)
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
// @Router /admin/ledger/icon/delete [post]
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
// @Router /admin/ledger/color/delete [post]
func (*LedgerAdminApi) DeleteColor(c *gin.Context) {
	id := c.Query("id")
	cd, err := ledgerAdminService.DeleteColorService(id)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(nil, cd, c)
}
