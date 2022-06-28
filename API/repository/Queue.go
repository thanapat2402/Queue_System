package repository

import (
	"fmt"
	"net/http"
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

	c.JSON(http.StatusOK, gin.H{"data": Queue, "message": "Created"})
}

func GetAllQueues(c *gin.Context) {
	queues := []model.QueueModel{}
	model.DB.Find(&queues)
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

func GetQueuesByType(c *gin.Context) {
	// Get model if exist
	queues := []model.QueueModel{}
	model.DB.Where("Type = ?", c.Param("Type")).Find(&queues)
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

func GetQueuesByCode(c *gin.Context) {
	// Get model if exist
	queues := []model.QueueModel{}
	if err := model.DB.Where("Code = ?", c.Param("Code")).First(&queues).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

func DeleterealQueue(c *gin.Context) {
	var queue model.QueueModel
	if err := model.DB.Where("Code = ?", c.Param("Code")).First(&queue).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	model.DB.Where("Code = ?", c.Param("Code")).Delete(&queue)

	c.JSON(http.StatusOK, gin.H{"data": queue, "message": "Deleted"})
}

func GenerateCode(genre string) (NewCode int) {
	queue := model.QueueModel{}
	model.DB.Where("Type=?", genre).Limit(1).Order("Date desc").Find(&queue)
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
