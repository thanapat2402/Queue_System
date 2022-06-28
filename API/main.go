package main

import (
	"q/model"

	"q/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Use(cors.Default())
	//connect to database + auto migrate
	model.ConnectDatabase()

	//Routes
	q := route.Group("/api/v1/queue")
	{
		q.GET("/", repository.GetAllQueues)
		q.GET("/:Type", repository.GetQueuesByType)
		q.GET("/code/:Code", repository.GetQueuesByCode)
		q.POST("/", repository.CreateQueue)
		q.DELETE("/:Code", repository.DeleteQueue)
	}

	//Run Server
	route.Run(":8086")
}
