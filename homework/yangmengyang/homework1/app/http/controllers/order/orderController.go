package order

import (
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/log/logger"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/http/controllers"
	"github.com/qit-team/snow-core/utils/httputil"
	"fmt"
)

func Servicea(c *gin.Context) {
	logger.Debug(c, "hello", "test message")
	fmt.Println("hello a")
	dict := map[string]interface{}{"param_a": "aaaa", "param_b": "bbb"}
	httputil.PostJson(c, "localhost:8080/service_b", dict)
	controllers.Success(c, "hello servicea!")
	return
}