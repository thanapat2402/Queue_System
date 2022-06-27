package repository

import (
	"fmt"
	"net/http"
	"q/handler"
	"q/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//----------------------------------------------------------------------------

func CreateQueue(c *gin.Context) {
	// Validate input
	var input model.QueueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Queue
	newCode := GenerateCode(input.Type)
	date := time.Now()
	code := fmt.Sprintf("%v%03d", input.Type, newCode)

	Queue := model.QueueModel{Code: code, Type: input.Type, Date: date}
	model.DB.Create(&Queue)

	c.JSON(http.StatusOK, gin.H{"data": Queue})
}

// func GetQueuesByType(genre string) {
// 	db, err := handler.DB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	tests := []model.QueueModel{}
// 	db.Where("Type=?", genre).Find(&tests)
// 	tests2 := new(model.QueueOp)
// 	for _, t := range tests {
// 		A := fmt.Sprintf("%v%03d", t.Type, t.Code)
// 		fmt.Printf("%v|%v|%v\n", A, t.Type, t.Date.Format("2006-02-01"))
// 		tests2.ResponseQueues(A, t.Type, t.Date)

// 		// fmt.Printf(`%v`, tests2)
// 		// return tests2

// 	}
// 	// c.JSON(http.StatusOK, gin.H{"data": tests2})
// }

func GetAllQueues(c *gin.Context) {
	tests := []model.QueueModel{}
	model.DB.Find(&tests)
	c.JSON(http.StatusOK, gin.H{"data": tests})
}

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
		strCode := strings.Trim(queue.Code, genre)
		intVar, _ := strconv.Atoi(strCode)
		NewCode := intVar + 1
		fmt.Println(NewCode)
		return NewCode
	}

	NewCode = 1
	fmt.Println(NewCode)
	return NewCode
}
