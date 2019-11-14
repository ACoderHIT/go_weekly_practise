package usermodel

import (
	"github.com/qit-team/snow-core/db"
	"sync"
)

var (
	once sync.Once
	m *userModel
)

type User struct {
	Id           int `xorm:"pk autoincr"` //注：使用getOne 或者ID() 需要设置主键
	Name         string
}

func (m *User) TableName() string {
	return "users"
}

type userModel struct {
	db.Model
}

func GetInstance() *userModel  {
	once.Do(func() {
		m = new(userModel)
	})
	return m
}

func (m *userModel) GetById(id int) (users []*User, err error) {
	users = make([]*User, 0)
	err = m.GetList(&users, "id = ?", []interface{}{id})
	return
}
