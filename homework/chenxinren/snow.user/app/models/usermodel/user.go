package usermodel

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/qit-team/snow-core/db"
	"os"
	"snow.user/app/utils"
	"sync"
)

var (
	once sync.Once
	m    *UserModel
)

//实体
type User struct {
	Id            int64     `xorm:"'id' bigint(20) pk autoincr"`
	Truename      string    `xorm:"'truename' varchar(20)"`
	Mobile        string    `xorm:"'mobile' varchar(11)"`
	CreatedAt	  string	`xorm:"'created_at'"`
	UpdatedAt	  string	`xorm:"'updated_at'"`
}

//表名
func (m *User) TableName() string {
	return "users"
}

//私有化，防止被外部new
type UserModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *UserModel {
	once.Do(func() {
		m = new(UserModel)
		//m.DiName = "" //设置数据库实例连接，默认db.SingletonMain
	})
	return m
}

func (m *UserModel) Insert(user *User) (id int64, err error) {
	return m.GetDb().Insert(user)
}

//获取ID
func (m *UserModel) GetIdByMobile(mobile string) (id int64, has bool, err error )  {
	var user = new(User)
	has, err = m.GetDb().Where("mobile = ?", mobile).Get(user) //不传user ===》 Table not found ?????
	id = user.Id
	fmt.Println(mobile, user, has)
	return
}

//批量注册mobile
func (m *UserModel) MulInsert(mobiles []string) (effected int64, err error) {

	f, err := os.Create("./logs/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	m.GetDb().SetLogger(xorm.NewSimpleLogger(f))
	m.GetDb().ShowSQL(true) //开启sql日志

	session := m.GetDb().NewSession()
	defer session.Clone()

	session.Begin() //开启事务

	users := make([]*User, len(mobiles))
	for i := 0; i < len(mobiles) ; i ++ {
		users[i] = new(User)
		users[i].Mobile = mobiles[i]
		users[i].CreatedAt = utils.NowTime()
		users[i].UpdatedAt = utils.NowTime()
	}
	effected, err = m.GetDb().Insert(users)
	fmt.Println(effected, err)

	if err != nil || effected != int64(len(mobiles)) {
		session.Rollback()
		return 0, errors.New("写入失败，已回滚")
	} else {
		session.Commit()
	}
	return
}

//更新姓名
func (m *UserModel) UpdateName(id int64, name string) (effected int64, err error) {
	user := new(User)
	user.Truename = name
	user.UpdatedAt = utils.NowTime()
	effected, err = m.GetDb().Id(id).Update(user)
	return
}

func (m *UserModel) UpdateMulName(id int64, name string) (effected int64, err error) {
	user := new(User)
	user.Truename = name
	user.UpdatedAt = utils.NowTime()
	effected, err = m.GetDb().Where("id < ?", id).Update(user)
	return
}

func (m *UserModel) UpdateMulNameBySql(id int64, name string) (res sql.Result, err error)  {
	return m.GetDb().Exec("update users set truename = ?, updated_at = now() where id < ?", name, id)
}