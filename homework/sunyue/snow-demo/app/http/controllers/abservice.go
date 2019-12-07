package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/utils/httputil"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/constants/errorcode"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/http/entities"
	"strconv"
	"time"
)

type responseData struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

var client httputil.Client

func init() {
	client = httputil.NewClient(time.Second * 5)
}

func HandleAGetDemoProcess(c *gin.Context)  {
	request := new(entities.ABServiceProcessRequest)
	err := GenRequest(c, request)
	fmt.Println(err)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}

	url := "http://ip:8080/week3/abservice/b/get?params=" + request.Params
	req, _ := httputil.NewGetRequest(url, nil)
	response, err := client.Do(context.TODO(), req)
	if err != nil {

		Error(c, errorcode.SystemError, "请求异常")
		return
	}
	result, err := httputil.DealResponse(response)

	resp := new(responseData)
	json.Unmarshal(result, resp)
	fmt.Print(resp)
	if resp.Code != 200 {
		Error(c, errorcode.SystemError, "请求结果异常")
		return
	}

	Success(c, resp.Data)
}

func HandleAPostDemoProcess(c *gin.Context)  {
	url := "http://10.11.3.69:8080/week3/abservice/b/post"
	req, _ := httputil.NewFormPostRequest(url, nil)
	response, err := client.Do(context.TODO(), req)
	if err != nil {

		Error(c, errorcode.SystemError, "请求异常")
		return
	}
	result, err := httputil.DealResponse(response)

	resp := new(responseData)
	json.Unmarshal(result, resp)
	fmt.Print(resp)
	if resp.Code != 200 {
		Error(c, errorcode.SystemError, "请求结果异常")
		return
	}

	Success(c, result)
}

func HandleAPostJsonDemoProcess(c *gin.Context)  {
	url := "http://10.11.3.69:8080/week3/abservice/b/post_json"
	req, _ := httputil.NewJsonPostRequest(url, nil)
	response, err := client.Do(context.TODO(), req)
	if err != nil {

		Error(c, errorcode.SystemError, "请求异常")
		return
	}
	result, err := httputil.DealResponse(response)

	resp := new(responseData)
	json.Unmarshal(result, resp)
	fmt.Print(resp.Data)
	if resp.Code != 200 {
		Error(c, errorcode.SystemError, "请求结果异常")
		return
	}

	Success(c, resp.Data)
}

func HandleBGetRequestProcess(c *gin.Context)  {
	params, err := strconv.Atoi(c.Query("params"))
	if err != nil {
		Error(c, errorcode.SystemError)
		return
	}
	retData := map[string]interface{}{
		"data" :params,
	}
	Success(c, retData)
}

func HandleBPostRequestProcess(c *gin.Context)  {
	Success(c, "success")
}

func HandleBPostJsonRequestProcess(c *gin.Context)  {
	Success(c, "success")
}
