package ucentercontroller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/controllers"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/entities"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/utils"
	"github.com/qit-team/snow-core/utils/httputil"
	_ "github.com/qit-team/snow-core/utils/httputil"
	"io/ioutil"
	"time"
)

const ServiceHost = "http://localhost:8080"

// HandleGetUserInfo godoc
// @Summary 获取用户资料
// @Description 获取用户资料
// @Tags snow
// @Accept  json
// @Produce json
// @Success 200 {array} entities.TestValidatorRequest
// @Failure 400 {object} controllers.HTTPError
// @Failure 404 {object} controllers.HTTPError
// @Failure 500 {object} controllers.HTTPError
// @Router /ucenter/user_info [get]
func HandleGetUserInfo(c *gin.Context) {
	userId := c.DefaultQuery("user_id", "")
	currentTime := time.Now()
	params := map[string]interface{}{"user_id": userId, "time": currentTime}
	getResp, err := httputil.Get(c, ServiceHost+"/user/info", params)
	if err != nil {
		controllers.Error(c, 500, err.Error())
		return
	}
	getContent, _ := ioutil.ReadAll(getResp.Body)

	postResp, err := httputil.Get(c, ServiceHost+"/user/info", params)
	if err != nil {
		controllers.Error(c, 500, err.Error())
		return
	}
	postContent, _ := ioutil.ReadAll(postResp.Body)

	postJsonResp, err := httputil.PostJson(c, ServiceHost+"/user/info", params)
	if err != nil {
		controllers.Error(c, 500, err.Error())
		return
	}
	postJsonContent, _ := ioutil.ReadAll(postJsonResp.Body)
	response := new(entities.ServiceResponse)
	json.Unmarshal(postJsonContent, response)

	if utils.TypeTransfer(response.Code) != "200" {
		controllers.Error(c, 500, "请求失败")
		return
	}

	content := map[string]interface{}{"get": getContent, "post": postContent, "postJson": postJsonContent}

	fmt.Printf("Results: %s\n", content)
	controllers.Success(c, content)
}
