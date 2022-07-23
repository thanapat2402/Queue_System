package handler

import (
	"fmt"
	"log"
	"os"
	"q/model"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"

	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// func ConnectDatabase() (db *gorm.DB) {

// 	//Set Data source name
// 	dsn := fmt.Sprintf("server=%v\\%v;Database=%v;praseTime=true",
// 		viper.GetString("db.server"),
// 		viper.GetString("db.driver"),
// 		viper.GetString("db.database"),
// 	)
// 	dial := sqlserver.Open(dsn)

// 	database, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}
// 	//auto migration
// 	database.AutoMigrate(&model.QueueModel{})
// 	return database
// }

func ConnectDatabase() (db *gorm.DB) {

	//Set Data source name
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	if os.Getenv("DB_DATABASE") == "" {
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?&parseTime=True&loc=Local",
			viper.GetString("db2.user"),
			viper.GetString("db2.pass"),
			viper.GetString("db2.host"),
			viper.GetString("db2.port"),
			viper.GetString("db2.database"),
		)
	}

	dial := mysql.Open(dsn)

	database, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("Failed to connect to database!")
	}
	//auto migration
	database.AutoMigrate(&model.QueueModel{})
	return database
}

func Readline() (secret string, token string) {
	if os.Getenv("CHANNEL_SECRET") == "" {
		secret := viper.GetString("line.CHANNEL_SECRET")
		token := viper.GetString("line.CHANNEL_TOKEN")
		return secret, token
	} else {
		secret := os.Getenv("CHANNEL_SECRET")
		token := os.Getenv("CHANNEL_TOKEN")
		return secret, token
	}
}

func GetBot() (bot *linebot.Client) {
	secret, token := Readline()
	bot, err := linebot.New(secret, token)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func initConfig() {
	//set Read form config.yaml
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
	//set timezone thailand
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func InitAll() {
	initTimeZone()
	initConfig()
}
