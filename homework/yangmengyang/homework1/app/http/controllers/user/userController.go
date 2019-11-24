package user

import (
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/log/logger"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/http/controllers"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/constants/errorcode"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/http/entities"
)

func Serviceb(c *gin.Context) {
	logger.Debug(c, "hello", "test messagebbb")
	request := new(entities.User)
	err := controllers.GenRequest(c, request)
	if err != nil {
		controllers.Error(c, errorcode.ParamError)
		return
	}
	logger.Debug(c, "params", request.Name, request.Mobile)
	controllers.Success(c, "hello service b!")
	return
}