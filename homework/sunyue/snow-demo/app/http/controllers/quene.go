package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/log/logger"
	"github.com/qit-team/snow-core/utils"
	"github.com/qit-team/work"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/constants/errorcode"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/http/entities"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/jobs/basejob"
	"fmt"
)

func HandleQueneProcess(c *gin.Context)  {
	request := new(entities.QueneProcessRequest)
	err := GenRequest(c, request)
	fmt.Println(err)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}

	logger.Info(c, "quene-test", "enter")

	jsonMessage, err := utils.JsonEncode(request)
	task := work.Task{
		Id: string(request.Id),
		Message: jsonMessage,
	}
	ok, err := basejob.EnqueueWithTask(c, "quene-test", task)
	if ok {
		Success(c, "success")
	} else {
		Error(c, errorcode.SystemError)
	}
}

