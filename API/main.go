package main

import (
	"q/handler"
	"q/model"
	"q/repository"
	"q/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	route := gin.Default()
	route.Use(cors.Default())
	//connect to database + auto migrate
	db := ConnectDatabase()

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

func ConnectDatabase() (db *gorm.DB) {

	//Set Data source name
	dsn := "server=localhost\\SQLEXPRESS;Database=QueueSystem;praseTime=true"
	dial := sqlserver.Open(dsn)

	database, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("Failed to connect to database!")
	}
	//auto migration
	database.AutoMigrate(&model.QueueModel{})
	return database
}
