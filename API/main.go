package main

import (
	"q/model"

	"q/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	model.ConnectDatabase()

	r.GET("/queues", repository.GetAllQueues)
	r.GET("/queues/:Type", repository.GetQueuesByType)
	r.GET("/queue/:Code", repository.GetQueuesByCode)
	r.POST("/queue", repository.CreateQueue)
	r.DELETE("/queue/:Code", repository.DeleterealQueue)

	r.Run(":8086")
}
