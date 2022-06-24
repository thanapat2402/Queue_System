package repository

import (
	"fmt"
	"q/handler"
	"q/model"
	"time"
)

func AutoMigrate() {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	err1 := db.AutoMigrate(model.QueueModel{})
	if err1 != nil {
		panic(err1)
	}
}

//----------------------------------------------------------------------------

func CreateQueue(genre string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	newCode := GenerateCode(genre)
	date := time.Now()
	test := model.QueueModel{Code: newCode, Type: genre, Date: date}
	db.Create(&test)
}

func GetQueuesByType(genre string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	tests := []model.QueueModel{}
	db.Where("Type=?", genre).Find(&tests)
	for _, t := range tests {
		// fmt.Printf("%03d|%v|%v\n", t.Code, t.Type, t.Date)
		fmt.Printf("%v%03d | %v\n", t.Type, t.Code, t.Date.Format("2006-02-01"))
	}
}

// func GetAllQueues() {
// 	db, err := handler.DB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	returns := make([]model.QueueOp, 0)
// 	tests := []model.QueueModel{}
// 	db.Find(&tests)
// 	var queue model.QueueOp
// 	for rows.Next() {
// 		if err := rows.Scan(&queue.Code, &queue.Type, &queue.Date); err != nil {
// 			return returns, err
// 		}
// 		returns = append(returns, queue)
// 	}
// 	fmt.Println(returns)
// 	return returns, nil
// 	// for _, t := range tests {

// 	// 	A := fmt.Sprintf("%v%03d", t.Type, t.Code)
// 	// 	fmt.Printf("%v|%v|%v\n", A, t.Type, t.Date.Format("2006-02-01"))
// 	// StrCode = append(StrCode, A, t.Type, t.Date)

// 	// fmt.Println("Code:",tests.StrCode)
// 	// fmt.Println(StrCode)

// }

func GetAllQueues1() {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	tests := []model.QueueModel{}
	db.Find(&tests)
	for _, t := range tests {

		A := fmt.Sprintf("%v%03d", t.Type, t.Code)
		fmt.Printf("%v|%v|%v\n", A, t.Type, t.Date.Format("2006-02-01"))
	}
}

// returncustomers := make([]model.CustomerModel, 0)
// execStr := fmt.Sprintf("exec SP_CustomerManagement @Process = N'%s'", "")
// rows, err := db.Query(execStr)
// if err != nil {
// 	return nil, err
// }
// var customer model.CustomerModel
// for rows.Next() {
// 	if err := rows.Scan(&customer.Id, &customer.CustomerId, &customer.FirstName, &customer.LastName); err != nil {
// 		return returncustomers, err
// 	}
// 	returncustomers = append(returncustomers, customer)
// }
// return returncustomers, nil

func DeleteTest(id uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	db.Delete(&model.QueueModel{}, id)
}

func DeleterealTest(id uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	db.Unscoped().Delete(&model.QueueModel{}, id)
}

func GenerateCode(genre string) (NewCode int) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	queue := model.QueueModel{}
	db.Where("Type=?", genre).Last(&queue)
	last := queue.Date.Format("2006-02-01")
	// fmt.Println(last)
	now := time.Now().Format("2006-02-01")
	// fmt.Println(now)

	if last == now {
		NewCode := queue.Code + 1
		fmt.Println(NewCode)
		return NewCode
	}

	NewCode = 1
	fmt.Println(NewCode)
	return NewCode
}
