package main

import (
	"q/handler"
	"q/repository"
	"q/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	web := route.Group("/api/v1/web")

	{
		web.GET("/", queueHandler.GetQueues)
		web.GET("/:Type", queueHandler.GetQueuesType)
		web.GET("/code/:Code", queueHandler.GetQueue)
		web.GET("/search", queueHandler.SearchQueue)
		web.POST("/", queueHandler.AddQueue)
		web.DELETE("/:Code", queueHandler.DeQueue)
	}

	line := route.Group("/api/v1/line")
	{
		line.GET("/", queueHandler.Hello)
		line.POST("/callback", queueHandler.Callback)
	}
	//Run Server
	route.Run()
}
