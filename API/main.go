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
	r.POST("/queues", repository.CreateQueue)

	r.Run(":8086")
}
