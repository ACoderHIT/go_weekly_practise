package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/http/ctxkit"
	"snow.user/app/constants/errorcode"
	"snow.user/app/http/entities"
	"snow.user/app/jobs/user_job"
	"snow.user/app/logic/user"
)

//入队操作
func EntryJob(c *gin.Context) {
	request := new(entities.JobValidatorRequest)
	err := GenRequest(c, request)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}
	jobParams := new(user_job.JobParams)
	jobParams.Mobile = request.Mobile
	jobParams.TraceId = ctxkit.GetTraceId(c)
	jsonParams, err := json.Marshal(jobParams)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}

	ok, err := user.DonEntry(c, string(jsonParams))
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}

	if !ok {
		Error(c, errorcode.SystemError)
		return
	}

	Success(c, "成功")
}
