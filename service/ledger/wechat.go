package ledger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/code"
	"server/config"
	"server/global"
	"server/models/ledger"
	"server/models/system"
	"server/utils"
)

type WeChatData struct {
	OpenID     string `json:"openId"`
	SessionKey string `json:"sessionKey"`
	UnionID    string `json:"unionId"`
	ErrCode    int    `json:"errCode"`
	ErrMsg     string `json:"errMsg"`
}

func (*LedgersService) GetOpenIDService(jsCode string) (data WeChatData, cd int, err error) {
	//获取openID
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", "wxd39978e0bb73fccb", "a879234494291e016bc7cb54fca06a87", jsCode)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return data, code.ErrorGetOpenID, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return data, code.ErrorGetOpenID, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, code.ErrorGetOpenID, err
	}
	return data, code.SUCCESS, nil
}
func (*LedgersService) GetTokenService(query ledger.WeChatUserInfo) (string, int, error) {
	db := global.DB
	//首先查询用户是否存在
	//根据openID查询用户
	var user system.User
	if err := db.Model(&user).Where("open_id = ?", query.OpenID).First(&system.User{}).Error; err != nil {
		//用户不存在
		//创建用户
		user := system.User{
			OpenID:   query.OpenID,
			UserName: query.Name,
			Avatar:   query.Avatar,
			NickName: query.Name,
		}
		if err := db.Create(&user).Error; err != nil {
			//创建token
			j := &utils.JWT{
				SigningKey: []byte(global.Config.JWT.SigningKey),
			}
			claims := j.CreateClaims(config.BaseClaims{
				ID:       user.ID,
				NickName: user.NickName,
				Username: user.UserName,
			})
			token, err := j.CreateToken(claims)
			if err != nil {
				//	颁发token错误
				return "", code.ErrorTokenGenerate, err
			}
			return token, code.SUCCESS, nil
		}
	}
	//用户存在
	//创建token
	j := &utils.JWT{
		SigningKey: []byte(global.Config.JWT.SigningKey),
	}
	claims := j.CreateClaims(config.BaseClaims{
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.UserName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		//	颁发token错误
		return "", code.ErrorTokenGenerate, err
	}
	return token, code.SUCCESS, nil
}
