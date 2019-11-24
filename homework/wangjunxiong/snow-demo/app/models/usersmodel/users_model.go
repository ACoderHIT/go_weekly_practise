package usersmodel

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/qit-team/snow-core/db"
	"strconv"
	"strings"
	"sync"
)

var (
	once sync.Once
	m    *UsersModel
)

const STATUS_CANCEL = -1
const STATUS_NORMAL = 0

//实体
type Users struct {
	AppId     int    `xorm:"'app_id' MEDIUMINT"`
	Email     string `xorm:"'email' varchar(250)"`
	Gender    string `xorm:"'gender' char(1)"`
	HeadIco   string `xorm:"'head_ico' varchar(250)"`
	Id        int64  `xorm:"'id' bigint(20) pk autoincr"`
	IdNumber  string `xorm:"'id_number' char(20)"`
	Mobile    string `xorm:"'mobile' varchar(11)"`
	Status    int    `xorm:"'status' tinyint(1)"`
	Password  string `xorm:"'password' varchar(64)"`
	Truename  string `xorm:"'truename' varchar(20)"`
	UserName  string `xorm:"'user_name' varchar(20)"`
	CreatedAt string `xorm:"'created_at' timestamp"`
	UpdatedAt string `xorm:"'updated_at' timestamp"`
}

//表名
func (m *Users) TableName() string {
	return "users"
}

//私有化，防止被外部new
type UsersModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *UsersModel {
	once.Do(func() {
		m = new(UsersModel)
		//m.DiName = "" //设置数据库实例连接，默认db.SingletonMain
	})
	return m
}

//更新用户信息

func (m *UsersModel) UpdateInfo(userId int, user Users) (int64, error) {
	rowsAffected, err := m.GetDb().Where("id = ?", user.Id).Update(user)
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//批量注册用户
func (m *UsersModel) MutilInsertUser(users []*Users) (rowsAffected int64, err error) {
	m.GetDb().ShowSQL(true)           //开启sql日志
	session := m.GetDb().NewSession() //事务
	rowsAffected, err = m.GetDb().Insert(users)
	if err != nil || rowsAffected != int64(len(users)) {
		session.Rollback()
		return 0, errors.New("操作失败")
	} else {
		session.Commit()
	}
	return
}

//批量恢复注销
func (m *UsersModel) MultiRecoverty() (rowsAffected int64, err error) {
	m.GetDb().ShowSQL(true) //开启sql日志
	user := new(Users)
	user.Status = STATUS_NORMAL
	rowsAffected, err = m.GetDb().Where("status = ?", STATUS_CANCEL).Update(user)
	if err != nil {
		return 0, errors.New("操作失败")
	}
	return
}

//更新用户状态
func (m *UsersModel) UpdateMultiStatusBySql(ids []int64, status int64) (res sql.Result, err error) {
	m.GetDb().ShowSQL(true) //开启sql日志
	var idStr []string
	var qStr []string
	for _, id := range ids {
		strId := strconv.FormatInt(id, 10)
		idStr = append(idStr, strId)
		qStr = append(qStr, "?")
	}
	statusStr := strconv.FormatInt(status, 10)
	query := fmt.Sprintf("update users set status = %s,updated_at = now() where id in (%s)", statusStr, strings.Join(idStr, ","))
	return m.GetDb().Exec(query)
}
