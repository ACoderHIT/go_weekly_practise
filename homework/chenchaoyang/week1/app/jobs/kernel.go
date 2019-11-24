package jobs

import (
	"github.com/qit-team/snow-core/log/logger"
	"github.com/qit-team/snow-core/queue"
	"github.com/qit-team/snow-core/redis"
	"github.com/qit-team/work"
	"go_weekly_practise/homework/chenchaoyang/week1/app/jobs/basejob"
	"go_weekly_practise/homework/chenchaoyang/week1/config"
	"strings"
)

/**
 * 配置队列任务
 */
func RegisterWorker(job *work.Job) {
	basejob.SetJob(job)

	//设置worker的任务投递回调函数
	job.AddFunc("topic-test", test)
	//设置worker的任务投递回调函数，和并发数
	job.AddFunc("topic-test1", test, 2)
	//使用worker结构进行注册
	job.AddWorker("topic-test2", &work.Worker{Call: work.MyWorkerFunc(test), MaxConcurrency: 1})

	RegisterQueueDriver(job)
	SetOptions(job)
}

/**
 * 给topic注册对应的队列服务
 */
func RegisterQueueDriver(job *work.Job) {
	//设置队列服务，需要实现work.Queue接口的方法
	q := queue.GetQueue(redis.SingletonMain, queue.DriverTypeRedis)
	//针对topic设置相关的queue
	job.AddQueue(q, "topic-test1", "topic-test2")
	//设置默认的queue, 没有设置过的topic会使用默认的queue
	job.AddQueue(q)
}

/**
 * 设置配置参数
 */
func SetOptions(job *work.Job) {
	//设置logger，需要实现work.Logger接口的方法
	job.SetLogger(logger.GetLogger())

	//设置启用的topic，未设置表示启用全部注册过topic
	if config.GetOptions().Queue != "" {
		topics := strings.Split(config.GetOptions().Queue, ",")
		job.SetEnableTopics(topics...)
	}
}
