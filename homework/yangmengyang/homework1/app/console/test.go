package console

import (
	"fmt"
	"time"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/models/order"
	"homework/go_weekly_practise/homework/yangmengyang/homework1/app/models/usermodel"
)

func test() {
	fmt.Println("run test")
}

type orderUser struct {
	order.Orders `xorm:"extends"`
	usermodel.Users `xorm:"extends"`
}

func yangtest() {
	//build/bin/snow -a command -m test
	fmt.Println("begin test")
	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD H:i:s11: ", currentTime.Format("2006-01-02 15:04:05"))
	// affected, err := engine.Insert(order)
	//insertSingle()

	//insertMult()

	//updateSingle()

	//updateMult()

	//joinTemp()

	useSession()

	// INSERT INTO user (name) values (?)
}

func insertSingle() {
	//插入单条数据
	orderStruct := order.Orders{}
	orderStruct.Name = "myname"
	orderStruct.OrderNo = "ymy0"
	orderStruct.CreatedAt = time.Now()
	affected, err := order.GetInstance().Insert(orderStruct)
	if err != nil {
		fmt.Println("error insert ", err)
	}
	fmt.Println("success11", affected)
}

func insertMult() {
	// 插入多条记录
	orders := make([]order.Orders, 3)
	orders[0].OrderNo = "ymy1"
	orders[1].OrderNo = "ymy2"
	orders[2].OrderNo = "ymy3"
	orders[0].CreatedAt = time.Now()
	orders[1].CreatedAt = time.Now()
	orders[2].CreatedAt = time.Now()
	affected, err := order.GetInstance().Insert(&orders)
	if err != nil {
		fmt.Println("error insert ", err)
	}
	fmt.Println("success11", affected)
}

func updateSingle()  {
	// 更新数据
	orderGet := order.Orders{}
	getResult, err2 := order.GetInstance().GetDb().Where("order_no = ?", "ymy0").Get(&orderGet)
	if err2 != nil {
		fmt.Println("error get ", err2)
	}
	fmt.Println("success22", getResult)
	updateId := orderGet.Id
	orderGet.Name = "Update"
	orderGet.Status = 0
	//affected3, err3 := order.GetInstance().Update(updateId, orderGet)
	affected3, err3 := order.GetInstance().GetDb().Id(updateId).Cols("name", "status").Update(orderGet)
	if err3 != nil {
		fmt.Println("error update ", err3)
	}
	fmt.Println("success333", affected3)
}

func updateMult()  {
	// 批量更新
	updateBatch := make([]order.Orders, 3)
	userId1 := 74437631
	// 使用原生代码
	err4 := order.GetInstance().GetDb().SQL("select * from orders where user_id = ? ", userId1).Find(&updateBatch)
	if err4 != nil {
		fmt.Println("error sql batch", err4)
	}
	fmt.Println("success444", updateBatch)
	updateBatch[0].Name = "qqq111"
	updateBatch[1].Name = "qqq222"
	updateBatch[2].Name = "qqq333"
	affect, err := order.GetInstance().GetDb().Update(updateBatch)
	if err != nil {
		fmt.Println("error update batch", err)
	}
	fmt.Println("success555", affect, updateBatch)
}

func joinTemp()  {
	users := make([]orderUser, 0)
	userId := 74437631
	usermodel.GetInstance().GetDb().Join("INNER", "orders", "users.id=orders.user_id").
		Where("users.id = ?", userId).Find(&users)
	fmt.Println("success666", users)
}

func useSession() {
	session := order.GetInstance().GetDb().NewSession()
	defer session.Close()
	err := session.Begin()
	order1 := order.Orders{}
	order1.Name = "myname1"
	order1.OrderNo = "ymyUseSession"
	order1.CreatedAt = time.Now()
	affected, err := session.Insert(&order1)
	if err != nil {
		session.Rollback()
		return
	}
	fmt.Println("useSession111 ", affected)

	// 更新数据
	order2 := order.Orders{Name: "ymyyyyy"}
	_, err1 := session.Where("id = ?", 989650472).Update(&order2)
	if err1 != nil {
		session.Rollback()
		return
	}

	_, err2 := session.Exec("delete from orders where id = ?", 989650474)
	if err2 != nil {
		session.Rollback()
		return
	}

	err3 := session.Commit()
	if err3 != nil {
		return
	}
}
