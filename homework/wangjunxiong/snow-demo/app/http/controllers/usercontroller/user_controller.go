package usercontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/controllers"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/http/entities/user"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/models/usersmodel"
	"github.com/go_weekly_practise-1/homework/wangjunxiong/snow-demo/app/services/userservice"
	"time"
)

func HandlePostUserLogin(c *gin.Context) {
	request := new(user.UserLoginRequest)
	request.LoginTime = time.Now()
	err := controllers.GenRequest(c, request)
	if err != nil {
		fmt.Println(err)
		controllers.Error(c, 500)
		return
	}
	insertId, err := userservice.InsertLoginInfo(request.UserId, request.Ip, request.LoginTime)

	if insertId > 0 {
		controllers.Success(c, 200)
	} else {
		fmt.Println(err);
		controllers.Error(c, 500)
	}
	return

}

func HandlUpdateUserInfo(c *gin.Context) {
	id := []int64{1, 2, 3}
	fmt.Println(usersmodel.GetInstance().UpdateMultiStatusBySql(id, 3))
	return
	request := new(user.UserRequest)
	err := controllers.GenRequest(c, request)
	if err != nil {
		fmt.Println(err)
		controllers.Error(c, 500, err.Error())
		return
	}
	afferctedRow, err := userservice.UpdateUserInfo(request.UserId, request.UpdateData)
	if afferctedRow > 0 {
		controllers.Success(c, 200)
	} else {
		fmt.Println(err);
		controllers.Error(c, 500, err.Error())
	}
	fmt.Println(request.UpdateData)
	return
}
