package processes

import (
	"context"
	"snow.user/app/jobs/basejob"
	"snow.user/app/jobs/user_job"
)

func Handle(c context.Context, jsonParams string) (ok bool, err error) {
	return basejob.Enqueue(c, user_job.QueueName, jsonParams)
}