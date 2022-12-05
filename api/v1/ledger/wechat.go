package ledger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/code"
)

func (*ApiLedger) GetOpenId(c *gin.Context) {
	jsCode := c.Query("js_code")
	fmt.Println(jsCode)
	ledgerService.GetOpenIDService(jsCode)

}

type WeChatUserInfo struct {
	//UnionID
	OpenID string `json:"open_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Phone  string `json:"phone"`
}

//GetToken 获取token
// @Tags WeChat
// @Summary 获取token
// @Description 获取token
// @Accept  json
// @Produce  json
// @Param data body WeChatUserInfo true "js_code"
// @Success 200 {object}  code.Response{data=string,code=int,msg=string,success=bool}
// @Router /ledger/wx/get_token [get]
func (*ApiLedger) GetToken(c *gin.Context) {
	var query WeChatUserInfo
	if err := c.ShouldBindJSON(&query); err != nil {
		fmt.Println(err)
	}
	token, cd, err := ledgerService.GetTokenService(query)
	if err != nil {
		code.FailResponse(cd, c)
		return
	}
	code.SuccessResponse(token, cd, c)
}
