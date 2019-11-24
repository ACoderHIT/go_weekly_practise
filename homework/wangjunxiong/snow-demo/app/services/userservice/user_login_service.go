package userservice

import (
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/models/userloginsmodel"
	"time"
)

//step3-1:单条数据写入，必须带有时间格式字段
func InsertLoginInfo(userId int, ip string, loginTime time.Time) (id int64, err error) {
	var instanceEntity userloginsmodel.UserLogins
	instanceEntity.UserId = userId
	instanceEntity.Ip = ip
	instanceEntity.LoginTime = loginTime
	id, err = userloginsmodel.GetInstance().InsertUserLogin(instanceEntity)
	return
}
