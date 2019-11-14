package user_job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qit-team/snow-core/log/logger"
	"github.com/qit-team/work"
	"snow.user/app/services/user"
)

const (
	QueueName = "queue-demo-cxr"
)

type JobParams struct {
	Mobile  string
	TraceId string
}

func Consumer(task work.Task) work.TaskResult {
	params := new(JobParams)
	err := json.Unmarshal([]byte(task.Message), params)
	if err != nil {
		logger.Error(context.Background(), QueueName, err.Error(), logger.NewWithField("task.Message", task.Message), logger.NewWithField("err", err))

		return work.TaskResult{
			Id:      task.Id,
			State:   work.StateFailed,
			Message: "",
		}
	}

	fmt.Println(task.Message)

	id, err := user.Create(params.Mobile)
	if err != nil {
		logger.Error(context.Background(), QueueName, err.Error(), logger.NewWithField("task.Message", task.Message), logger.NewWithField("err", err))
	}

	return work.TaskResult{
		Id:      task.Id,
		State:   work.StateSucceed,
		Message: "Id:" + string(id),
	}
}
