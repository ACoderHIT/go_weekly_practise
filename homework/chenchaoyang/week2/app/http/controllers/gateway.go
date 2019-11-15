package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/utils/httputil"
	"go_weekly_practise/homework/chenchaoyang/week2/app/http/entities"
	"go_weekly_practise/homework/chenchaoyang/week2/app/utils"
	"io/ioutil"
)

const (
	ServiceHost = "http://127.0.0.1:8080"
)

func GatewayGetRequest(c *gin.Context) {
	value := c.Query("value")
	params := map[string]interface{}{"from": "gateway", "value": value}
	resp, err := httputil.Get(c, ServiceHost+"/service/dealGet", params)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Results: %s\n", content)

	Success(c, string(content))
}

func GatewayPostRequest(c *gin.Context) {
	value := c.Query("value")
	params := map[string]interface{}{"from": "gateway", "value": value}
	resp, err := httputil.Post(c, ServiceHost+"/service/dealPost", params)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Results: %s\n", content)

	Success(c, string(content))
}

func GatewayPostJsonRequest(c *gin.Context) {
	value := c.Query("value")
	params := map[string]interface{}{"from": "gateway", "value": value}
	resp, err := httputil.Post(c, ServiceHost+"/service/dealPost", params)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	response := new(entities.ServiceResponse)
	json.Unmarshal(content, response)
	//请求失败
	if utils.TypeTransfer(response.Code) != "200" {
		Error(c, 500, "service code error")
		return
	}
	//请求成功
	Success(c, response.Data)
}
