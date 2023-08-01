package halo

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/netutil"
	"io/ioutil"
	"net/http"
	code2 "server/code"
	"server/models/halo"
	"strconv"
)

type ServiceHalo struct{}

func (*ServiceHalo) SaveArticle(query halo.ArticleHaloResponse) (token interface{}, code int, err error) {
	url := query.Url + "/api/admin/posts"
	header := map[string]string{
		"Content-Type":        "application/json",
		"Admin-Authorization": query.Token,
	}
	bodyParams, _ := json.Marshal(query)

	var resp *http.Response
	if query.ID != 0 {
		url = url + "/" + strconv.FormatInt(query.ID, 10)
		resp, err = netutil.HttpPut(url, header, nil, bodyParams)
	} else {
		resp, err = netutil.HttpPost(url, header, nil, bodyParams)
	}
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
