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

//Create new Queue
func CreateQueue(c *gin.Context) {
	// Validate input
	var input model.QueueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCode := GenerateCode(input.Type)
	date := time.Now()
	code := fmt.Sprintf("%v%03d", input.Type, newCode)
	// Create Queue
	Queue := model.QueueModel{Code: code, Type: input.Type, Date: date, Name: input.Name, Tel: input.Tel}
	model.DB.Create(&Queue)

	c.JSON(http.StatusOK, gin.H{"data": Queue, "message": "Created"})
}

//Get All Queues
func GetAllQueues(c *gin.Context) {
	queues := []model.QueueModel{}
	model.DB.Order("Date").Find(&queues)
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

//Get by Type
func GetQueuesByType(c *gin.Context) {
	queues := []model.QueueModel{}
	model.DB.Where("Type = ?", c.Param("Type")).Order("Date").Find(&queues)
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

//Get a queue by code
func GetQueuesByCode(c *gin.Context) {
	queues := []model.QueueModel{}
	if err := model.DB.Where("Code = ?", c.Param("Code")).First(&queues).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

//Delete a queue
func DeleteQueue(c *gin.Context) {
	// Get model if exist
	var queue model.QueueModel
	if err := model.DB.Where("Code = ?", c.Param("Code")).First(&queue).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	model.DB.Where("Code = ?", c.Param("Code")).Delete(&queue)

	c.JSON(http.StatusOK, gin.H{"data": queue, "message": "Deleted"})
}

//Gen Runningnumber
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
