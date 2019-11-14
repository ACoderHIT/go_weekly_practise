package controllers

import (
	"github.com/gin-gonic/gin"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/constants/errorcode"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/http/entities"
	"fmt"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/services/userservice"
)

func HandleUserQueryProcess(c *gin.Context)  {
	request := new(entities.UserQueryRequest)
	err := GenRequest(c, request)
	fmt.Println(err)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}

	id := request.Id

	list, err := userservice.GetUserById(id)
	if err != nil {
		Error(c, errorcode.SystemError)
	}

	Success(c, list)
}
