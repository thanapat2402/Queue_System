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
	r.POST("/queues", repository.CreateQueue)
	r.DELETE("/queues/:Code", repository.DeleterealQueue)

	r.Run(":8086")
}
