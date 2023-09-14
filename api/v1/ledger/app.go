package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/code"
)

type UpdateAppStruct struct {
	AppKey string `json:"app_key"`
	ApiKey string `json:"_api_key"`
}

// GetAppNeedUpdate 获取app是否需要更新
// @Tags statistics
// @Summary 获取app是否需要更新
// @Description 获取app是否需要更新
// @Accept  json
// @Produce  json
// @Success 200 {object}  code.Response{}
// @Router /ledger/app/need_update [get]
func (*ApiLedger) GetAppNeedUpdate(c *gin.Context) {
	url := "https://www.pgyer.com/apiv2/app/check?_api_key=c3b11d1ce00808553eeed3276d33e386&appKey=62a36ed833fa3a0b0749a900f7b715da"
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		code.FailResponse(code.ErrorGetAppNeedUpdate, c)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		code.FailResponse(code.ErrorGetAppNeedUpdate, c)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		code.FailResponse(code.ErrorGetAppNeedUpdate, c)
		return
	}
	type AppResponse struct {
		Code int `json:"code"`
		Data struct {
			BuildBuildVersion      string `json:"buildBuildVersion"`
			DownloadURL            string `json:"downloadURL"`
			BuildVersion           string `json:"buildVersion"`
			BuildUpdateDescription string `json:"buildUpdateDescription"`
		}
	}
	var appResponse AppResponse
	err = json.Unmarshal(body, &appResponse)
	if err != nil {
		return
	}
	code.SuccessResponse(appResponse.Data, code.SUCCESS, c)
}
