package main

import (
	"fmt"
	"q/handler"
	"q/model"
	"q/repository"
	"q/service"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	initTimeZone()
	initConfig()
	route := gin.Default()
	route.Use(cors.Default())
	//connect to database + auto migrate
	db := ConnectDatabase()

	queueRepo := repository.NewQueueRepositoryDB(db)
	// queueRepo := repository.NewQueueRepositoryMock()
	queueService := service.NewQueueService(queueRepo)
	queueHandler := handler.NewQueueHandler(queueService)

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

func ConnectDatabase() (db *gorm.DB) {

	//Set Data source name
	dsn := fmt.Sprintf("server=%v\\%v;Database=%v;praseTime=true",
		viper.GetString("db.server"),
		viper.GetString("db.driver"),
		viper.GetString("db.database"),
	)
	dial := sqlserver.Open(dsn)

	database, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("Failed to connect to database!")
	}
	//auto migration
	database.AutoMigrate(&model.QueueModel{})
	return database
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
