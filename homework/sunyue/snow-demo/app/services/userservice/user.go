package userservice

import (
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/models/usermodel"
	"os"
)

func GetUserById(id int) (result sql.Result, err error) {

	file, err := os.Create("/Users/sunyue/go/src/go_weekly_practise-1/homework/sunyue/snow-demo/logs/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	db := usermodel.GetInstance().GetDb()

	db.SetLogger(xorm.NewSimpleLogger(file))
	db.ShowSQL(true) //开启sql日志

	result, err = db.Exec("select * from users left join user_addresses on users.id=user_addresses.user_id where users.id = ?", id)
	fmt.Print("users is", result)
	return
}

type User struct {
	Name  string
	Id int
}
