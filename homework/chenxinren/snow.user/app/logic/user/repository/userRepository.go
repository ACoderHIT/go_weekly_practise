package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qit-team/snow-core/log/logger"
	"snow.user/app/utils/httputil"
	"time"
)

type responseData struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func GetUserById(c context.Context, user_id int64)  {
	client := httputil.NewClient(time.Second * 5)
	url := "http://127.0.0.1:8080/hello"
	//req, _ := httputil.NewGetRequest(url, nil)
	//req, _ := httputil.NewPostRequest(url, nil)

	params := make(map[string]interface{})
	params["user_id"] = user_id

	req, _ := httputil.NewJsonPostRequest(url, params)
	response, err := client.Do(c, req)
	if err != nil {
		logger.Error(c, err.Error())
	}
	result, err := httputil.DealResponse(response)
	fmt.Println(string(result))

	resp := new(responseData)
	json.Unmarshal(result, resp)
	if resp.Code != 200 {
		logger.Error(c, "get result is not ok")
	}
}
