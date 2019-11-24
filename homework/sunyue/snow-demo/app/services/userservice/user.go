package userservice

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/models/usermodel"
	"os"
)

func GetUserById(id int) (users []map[string]byte, err error) {

	file, err := os.Create("/Users/sunyue/go/src/go_weekly_practise/homework/sunyue/snow-demo/logs/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	usermodel.GetInstance().GetDb().SetLogger(xorm.NewSimpleLogger(file))
	usermodel.GetInstance().GetDb().ShowSQL(true) //开启sql日志

	var sliceOfStructs []User
	usermodel.GetInstance().GetDb().Join("LEFT", "user_addresses", "users.id = user_addresses.user_id").Where("users.id = ?", id).Find(&sliceOfStructs)

	fmt.Print("user by id id ", sliceOfStructs)
	return
}

type User struct {
	Name  string
	Id int
}
