package main

// import (
// 	"fmt"
// 	"q/repository"
// 	"q/service"
// )

import (
	"q/handler"
	"q/repository"
	"q/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Use(cors.Default())
	//connect to database + auto migrate
	db := repository.ConnectDatabase()

	queueRepo := repository.NewQueueRepositoryDB(db)
	queueService := service.NewQueueService(queueRepo)
	queueHandler := handler.NewQueueHandler(queueService)

	//Routes
	q := route.Group("/api/v1/queue")

	{
		q.GET("/", queueHandler.GetQueues)
		q.GET("/:Type", queueHandler.GetQueuesType)
		q.GET("/code/:Code", queueHandler.GetQueue)
		q.POST("/", queueHandler.AddQueue)
		q.DELETE("/:Code", queueHandler.DeQueue)
	}

	//Run Server
	route.Run(":8086")
}

// func main() {
// 	db := repository.ConnectDatabase()

// 	queueRepo := repository.NewQueueRepositoryDB(db)
// 	queueService := service.NewQueueService(queueRepo)

// 	queues, err := queueService.GetQueues()
// 	// queues, err := queueService.GetQueue("A011")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(queues)
// 	// All,_ := queueRepo.GetAllQueues()
// 	// fmt.Println(All)

// }
