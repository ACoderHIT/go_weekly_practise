package order

import (
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/log/logger"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/http/controllers"
	"fmt"
	"github.com/qit-team/snow-core/utils/httputil"
)

func Servicea(c *gin.Context) {
	logger.Debug(c, "hello", "test message")
	fmt.Println("hello a")
	postJsonHw(c)
	fmt.Println("goodby a")
	controllers.Success(c, "hello servicea!")
	return
}

func postJsonHw(c *gin.Context) {
	dict := map[string]interface{}{"name": "aaaa", "mobile": "12345678901"}
	result, err := httputil.PostJson(c, "http://localhost:8080/service_b", dict)
	if err != nil {
		fmt.Println("exception", err)
	}
	fmt.Println("success", result)
}