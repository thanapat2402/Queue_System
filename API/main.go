package main

import (
	"fmt"
	"q/handler"
	"q/repository"
	"q/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	handler.InitAll()
	//connect to database + auto migrate
	db := handler.ConnectDatabase()

	//Use Mock Data Repository to test
	// queueRepo := repository.NewQueueRepositoryMock2()

	queueRepo := repository.NewQueueRepositoryDB(db)
	queueService := service.NewQueueService(queueRepo)
	queueHandler := handler.NewQueueHandler(queueService)

	route := gin.Default()
	route.Use(cors.Default())
	//Routes
	q := route.Group("/api/v1/queue")

	{
		q.GET("/", queueHandler.GetQueues)
		q.GET("/:Type", queueHandler.GetQueuesType)
		q.GET("/code/:Code", queueHandler.GetQueue)
		q.GET("/search", queueHandler.SearchQueue)
		q.POST("/", queueHandler.AddQueue)
		q.DELETE("/:Code", queueHandler.DeQueue)
	}

	//Run Server
	route.Run(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
