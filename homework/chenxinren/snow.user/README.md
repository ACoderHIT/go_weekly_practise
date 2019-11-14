## Snow
Snow是一套简单易用的Go语言业务框架，整体逻辑设计简洁，支持HTTP服务、队列调度和任务调度等常用业务场景模式。

## Quick start

### Build
sh build/shell/build.sh

### Run
```shell
1. build/bin/snow -a api  #启动Api服务
2. build/bin/snow -a cron #启动Cron定时任务服务
3. build/bin/snow -a job  #启动队列调度服务
4. build/bin/snow -a command -m test  #执行名称为test的脚本任务
```

## Documents

- [项目地址](https://github.com/qit-team/snow)
- [中文文档](https://github.com/qit-team/snow/wiki)
- [changelog](https://github.com/qit-team/snow/blob/master/CHANGLOG.md)
- [xorm](http://gobook.io/read/github.com/go-xorm/manual-zh-CN/)
