package jobs

import (
	"encoding/json"
	"github.com/qit-team/snow-core/log/logger"
	"github.com/qit-team/work"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/services/queneservice"
	"log"
	"time"
	"context"
	"fmt"
)

func quene(task work.Task) (work.TaskResult) {
	time.Sleep(time.Millisecond * 5)
	s, err := work.JsonEncode(task)
	messageStr := task.Message
	fmt.Printf("%s", messageStr)
	message := Message{}
	errMessage := json.Unmarshal([]byte(messageStr), &message)
	if errMessage != nil {
		log.Fatal(err)
	}
	num := message.Num
	update := message.Update
	id := message.Id
	if err != nil {
		//work.StateFailed 不会进行ack确认
		//work.StateFailedWithAck 会进行actk确认
		//return work.TaskResult{Id: task.Id, State: work.StateFailed}
		return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
	} else {
		//work.StateSucceed 会进行ack确认
		logger.Info(context.TODO(), "quene-test", s, "exit")
		logId, errInsert := queneservice.InsertQueneLog(s, id, update, num)
		if errInsert != nil {
			logger.Error(context.TODO(), "quene-test", "insert err")
			fmt.Println(errInsert)
		}
		return work.TaskResult{Id: string(logId), State: work.StateSucceed}
	}

}

type Message struct {
	Name  string
	Id int
	Update bool
	Num int
}