package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	code2 "server/code"
	"server/models/system"
)

type SkiService struct{}

func (*SkiService) UploadFile(fileHeader *multipart.FileHeader, file multipart.File, config system.UserConfig) (string, string, int, error) {
	url := "https://img.ski/api/v1/upload"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", fileHeader.Filename)
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return "", "", code2.ErrorOpenFile, errFile1
	}
	err := writer.Close()
	if err != nil {
		return "", "", code2.ErrorOpenFile, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", "", code2.ErrorSkiUpload, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", "", code2.ErrorSkiUpload, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", "", code2.ErrorReadThirdPartyResponse, err
	}
	type SkiStruct struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Key   string `json:"key"`
			Name  string `json:"name"`
			Size  string `json:"size"`
			Links struct {
				Url string `json:"url"`
			}
		}
	}
	var skiData SkiStruct
	err = json.Unmarshal(body, &skiData)
	if !skiData.Status {
		return "", "", code2.ErrorSkiUpload, fmt.Errorf(skiData.Message)
	}
	return skiData.Data.Links.Url, skiData.Data.Key, 0, nil
}

// DeleteFile 删除文件
func (*SkiService) DeleteFile(key string, token string) error {
	url := "https://img.ski/api/v1/images/" + key
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	type SkiStruct struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	var skiData SkiStruct
	err = json.Unmarshal(body, &skiData)
	if !skiData.Status {
		return fmt.Errorf(skiData.Message)
	}
	return nil
}
