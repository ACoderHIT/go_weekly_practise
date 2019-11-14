package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/log/logger"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/http/controllers"
)

func Serviceb(c *gin.Context) {
	logger.Debug(c, "hello", "test message")
	controllers.Success(c, "hello service b!")
	return
}