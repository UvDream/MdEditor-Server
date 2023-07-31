package system

import (
	"github.com/gin-gonic/gin"
	"server/code"
	"server/models/system"
	"server/utils"
)

// Feedback 建议反馈
// @Tags system
// @Summary 建议反馈
// @Produce  json
// @Param  feedback body system.Feedback true "建议反馈"
// @Success 200 {string} code.Response{"success":true,"data":string,"msg":"设置成功"}
// @Router /user/feedback [post]
func (*UserApi) Feedback(c *gin.Context) {
	var feedback system.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		code.FailResponse(code.ErrorSetUserConfigMissingParam, c)
		return
	}
	feedback.UserID = utils.FindUserID(c)
	data, cd, err := userService.FeedbackService(feedback)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(data, cd, c)
}

// FeedbackList 建议反馈列表
// @Tags system
// @Summary 建议反馈列表
// @Accept  json
// @Produce  json
// @Param  query  query    models.PaginationRequest  true  "参数"
// @Param  query  query    system.Feedback true  "参数"
// @Success 200 {object} code.Response{code=int,msg=string,success=bool,data=[]system.Feedback}
// @Router /user/feedback_list [get]
func (*UserApi) FeedbackList(c *gin.Context) {
	//获取参数
	var query system.Feedback
	if err := c.ShouldBindQuery(&query); err != nil {
		code.FailResponse(code.ErrorMissingLedgerId, c)
		return
	}
	keyword := c.Query("key_word")
	// 获取用户id
	userId := utils.FindUserID(c)
	data, total, cd, err := userService.FeedbackListService(userId, query, keyword, c)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponseList(data, total, cd, c)
}
