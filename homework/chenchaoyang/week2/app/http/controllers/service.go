package controllers

import (
	"github.com/gin-gonic/gin"
	"go_weekly_practise/homework/chenchaoyang/week2/app/http/entities"
	"go_weekly_practise/homework/chenchaoyang/week2/app/utils"
)

func DealGetRequest(c *gin.Context) {
	value := c.Query("value")
	Success(c, "service receive from gateway:"+value)
}

func DealPostRequest(c *gin.Context) {
	value := c.PostForm("value")
	Success(c, "service receive from gateway:"+value)
}

func DealPostJsonRequest(c *gin.Context) {
	request := new(entities.CustomValidatorRequest)
	err := GenRequest(c, request)
	if err != nil {
		Error(c, 500, "parse param error:"+err.Error())
		return
	}
	value := utils.TypeTransfer(request.Value)
	Success(c, "service receive from gateway:"+value)
}
