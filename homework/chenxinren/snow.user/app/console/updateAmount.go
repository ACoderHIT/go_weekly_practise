package console

import (
	"context"
	"github.com/qit-team/snow-core/log/logger"
	"snow.user/app/services/user"
	"time"
)

type Params struct {
	Mobile     string
}

func insert()  {
	//问题，如何在command获取自定义的参数
	param := new(Params)
	param.Mobile = string(time.Now().Unix())
	id, err := user.Create(param.Mobile)
	if err != nil {
		logger.Error(context.Background(), "updateAmount", err.Error(), logger.NewWithField("err", err))
	}
	logger.Info(context.Background(), "id", id)
}