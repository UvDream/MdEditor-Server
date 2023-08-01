package halo

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/netutil"
	"io/ioutil"
	code2 "server/code"
	"server/models/halo"
)

func (*ServiceHalo) GetToken(query halo.UserHalo) (token interface{}, code int, err error) {
	url := query.Url + "/api/admin/login"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	bodyParams, _ := json.Marshal(query)
	resp, err := netutil.HttpPost(url, header, nil, bodyParams)
	if err != nil {
		return "", code2.HaloServerError, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//从body拿出数据
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", code2.HaloServerError, err
	}
	return result["data"], code2.SUCCESS, nil
}
