package userservice

import (
	"fmt"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/models/usersmodel"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/utils"
)

//step3-2:单条数据写入，必须带有时间格式字段
func InserUserBatch(mobils []string) (rowsAffected int64, err error) {
	users := make([]*usersmodel.Users, len(mobils))
	for i := 0; i < len(mobils); i++ {
		tmpUser := new(usersmodel.Users)
		tmpUser.Mobile = mobils[i]
		tmpUser.CreatedAt = utils.NowTime()
		users[i] = tmpUser
	}
	fmt.Println(users);
	rowsAffected,err = usersmodel.GetInstance().MutilInsertUser(users)
	return
}

//step3-3:单条数据更新，注意0值和非0值
func UpdateUserInfo(userId int, updateData map[string]interface{}) (rowsAffected int64, err error) {
	var user usersmodel.Users
	updSign := false
	_, err = usersmodel.GetInstance().GetOne(userId, &user)
	fmt.Println(user)
	if err != nil {
		return
	}
	if _, ok := updateData["password"]; ok {
		user.Password = updateData["password"].(string)
		updSign = true
	}
	if updSign {
		rowsAffected, err = usersmodel.GetInstance().UpdateInfo(userId, user)
	} else {
		fmt.Println("without update")

		rowsAffected = 1
	}
	return
}
