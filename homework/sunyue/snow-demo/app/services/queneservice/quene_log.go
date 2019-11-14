package queneservice

import (
	"database/sql"
	"github.com/go-xorm/xorm"
	"go_weekly_practise-1/homework/sunyue/snow-demo/app/models/quenemodel"
	"os"
	"time"
	"fmt"
)

func InsertQueneLog(queneData string, id int, update bool, num int) (affected int64, err error)  {
	if update {
		updateTable(queneData, id)
	} else {
		insertIntoTable(queneData, num)
	}
	return
}

func insertIntoTable(queneData string, num int) (affected int64, err error) {
	queneLog := quenemodel.QueneLog{}
	queneLogInsert := make([]*quenemodel.QueneLog, num)
	for i := 0; i < num; i ++ {
		queneLogInsert[i] = new(quenemodel.QueneLog)
		queneLogInsert[i].Data = queneData;
		queneLogInsert[i].Update_time = time.Now()
	}
	fmt.Printf("%s", queneLogInsert)

	file, err := os.Create("/Users/sunyue/go/src/go_weekly_practise/homework/sunyue/snow-demo/logs/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	quenemodel.GetInstance().GetDb().SetLogger(xorm.NewSimpleLogger(file))
	quenemodel.GetInstance().GetDb().ShowSQL(true) //开启sql日志

	session := quenemodel.GetInstance().GetDb().NewSession()
	defer session.Clone()

	session.Begin()

	affected, err = quenemodel.GetInstance().GetDb().Table(queneLog.TableName()).Insert(queneLogInsert)

	err = session.Commit()

	if err != nil{
		session.Rollback()
		panic(err.Error())
	}
	return
}

func updateTable(queneData string, id int) (affected int64, err error, result sql.Result) {
	queneLog := quenemodel.QueneLog{}

	queneLogUpdate := new(quenemodel.QueneLog)
	queneLogUpdate.Data = queneData;
	queneLogUpdate.Update_time = time.Now()

	affected, err = quenemodel.GetInstance().GetDb().Table(queneLog.TableName()).Where("id < ?", id).Update(queneLogUpdate)
	result, err = quenemodel.GetInstance().GetDb().Exec("Update test_go set update_time = now() where id = ?", id)

	return
}