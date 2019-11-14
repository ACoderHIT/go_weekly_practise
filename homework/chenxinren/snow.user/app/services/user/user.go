package user

import (
	"context"
	"github.com/qit-team/snow-core/log/logger"
	"snow.user/app/models/usermodel"
	"time"
)

func Create(mobile string) (id int64, err error) {
	um := usermodel.GetInstance()
	user := &usermodel.User{}
	user.Mobile = mobile
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	id, err = um.Insert(user)
	return
}

func MulCreate(mobiles []string) (int64, error)  {
	um := usermodel.GetInstance()
	effected, err := um.MulInsert(mobiles)
	if err != nil {
		logger.Error(context.Background(), "err", err)
		return 0, err
	}

	return effected, nil
}

func UpdateNameByMobile(mobile string, name string) (int64, error) {

	um := usermodel.GetInstance()
	id, has, err := um.GetIdByMobile(mobile)
	if err != nil {
		return 0, err
	}

	if has != true {
		return 0, nil
	}

	return um.UpdateName(id, name)
}

func UpdateMulNameById(id int64, name string) (int64, error) {

	um := usermodel.GetInstance()
	return um.UpdateMulName(id, name)
}

