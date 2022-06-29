package repository

import (
	"q/model"
)

//Port
type QueueRepository interface {
	GetAllQueues() ([]model.QueueModel, error)
	GetQueuesByType(Type string) ([]model.QueueModel, error)
	GetQueuesByCode(Code string) (*model.QueueModel, error)
	CreateQueue(data model.QueueInput) (*model.QueueModel, error)
	DeleteQueue(Code string) (*model.QueueModel, error)
}

// //Create new Queue
// func CreateQueue(c *gin.Context) {
// 	// Validate input
// 	var input model.QueueInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	newCode := generateCode(input.Type)
// 	date := time.Now()
// 	code := fmt.Sprintf("%v%03d", input.Type, newCode)
// 	// Create Queue
// 	Queue := model.QueueModel{Code: code, Type: input.Type, Date: date, Name: input.Name, Tel: input.Tel}
// 	model.DB.Create(&Queue)

// 	c.JSON(http.StatusOK, gin.H{"data": Queue, "message": "Created"})
// }

//Get All Queues
// func GetAllQueues(c *gin.Context) {
// 	queues := []model.QueueModel{}
// 	model.DB.Order("Date").Find(&queues)
// 	c.JSON(http.StatusOK, gin.H{"data": queues})
// }

//Get by Type
// func GetQueuesByType(c *gin.Context) {
// 	queues := []model.QueueModel{}
// 	model.DB.Where("Type = ?", c.Param("Type")).Order("Date").Find(&queues)
// 	c.JSON(http.StatusOK, gin.H{"data": queues})
// }

// //Get a queue by code
// func GetQueuesByCode(c *gin.Context) {
// 	queues := []model.QueueModel{}
// 	if err := model.DB.Where("Code = ?", c.Param("Code")).First(&queues).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": queues})
// }

// //Delete a queue
// func DeleteQueue(c *gin.Context) {
// 	// Get model if exist
// 	var queue model.QueueModel
// 	if err := model.DB.Where("Code = ?", c.Param("Code")).First(&queue).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	model.DB.Where("Code = ?", c.Param("Code")).Delete(&queue)

// 	c.JSON(http.StatusOK, gin.H{"data": queue, "message": "Deleted"})
// }
