package console

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"go_weekly_practise/homework/chenchaoyang/week1/app/models/homeworkmodel"
	"os"
	"strconv"
)

func insertOne() {
	var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(os.Stdout)
	homeworkmodel.GetInstance().GetDb().Engine.SetLogger(logger)
	homeworkmodel.GetInstance().GetDb().Engine.ShowSQL(true)

	order := homeworkmodel.Orders{
		UserId:  "74446023",
		OrderNo: "2222222",
		Status:  10,
	}
	//如何取插入的id
	affected, err := homeworkmodel.GetInstance().Insert(order)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert one success:" + strconv.FormatInt(affected, 10))
}

func InsertBatch() {
	orders := make([]homeworkmodel.Orders, 2)
	for index := 0; index < len(orders); index++ {
		orders[index].OrderNo = "ccy_order_" + strconv.Itoa(index)
		orders[index].UserId = "74446023"
	}

	affected, err := homeworkmodel.GetInstance().Insert(orders)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert one success:" + strconv.FormatInt(affected, 10))
}

func UpdateOne() {
	order := homeworkmodel.Orders{
		UserId:  "74446023",
		OrderNo: "2222222",
		Status:  0, //0值
	}
	affected, err := homeworkmodel.GetInstance().Update(989650473, order, "status")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert one success:" + strconv.FormatInt(affected, 10))
}

func UpdateBatch() {
	order := homeworkmodel.Orders{
		Status: 100,
	}
	affected, err := homeworkmodel.GetInstance().GetDb().Where("user_id = ?", "74446023").Update(order)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert one success:" + strconv.FormatInt(affected, 10))
}

func SelectJoin() {
	orders := make([]*homeworkmodel.Orders, 0)
	err := homeworkmodel.GetInstance().GetDb().Join(
		"INNER",
		"order_devices",
		"orders.id=order_devices.order_id").Where("orders.user_id = ?", "74446023").Find(&orders)
	if err != nil {
		fmt.Println(err.Error())
	}
	for index := range orders {
		fmt.Println(orders[index])
	}
}

func UseSession() {
	session := homeworkmodel.GetInstance().GetDb().NewSession()
	defer session.Close()
	err := session.Begin()
	order := homeworkmodel.Orders{
		UserId:  "74446023",
		OrderNo: "111111",
		Status:  10,
	}
	_, err = session.Insert(&order)
	if err != nil {
		session.Rollback()
		fmt.Println("rollback")
		return
	}

	err = session.Commit()
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func UseAnotherDatabase() {

}

func ShowSqlLog() {
	var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(os.Stdout)
	homeworkmodel.GetInstance().GetDb().Engine.SetLogger(logger)
	homeworkmodel.GetInstance().GetDb().Engine.ShowSQL(true)

	var order homeworkmodel.Orders
	homeworkmodel.GetInstance().GetDb().Where("user_id = ?", "74446023").Get(&order)
	fmt.Println(order)
}

func UseSql() {
	results, err := homeworkmodel.GetInstance().GetDb().Engine.Query("select * from orders where user_id = 74446023 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(results)
}
