package model

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {

	//Set Data source name
	dsn := "server=localhost\\MSSQLSERVER;Database=QueueSystem;praseTime=true"
	dial := sqlserver.Open(dsn)

	database, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&QueueModel{})

	DB = database
}
