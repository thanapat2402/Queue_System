package main

import (
	"q/repository"
	// 	"q/controller"
	// 	"github.com/gin-contrib/cors"
	// 	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()
// 	r.Use(cors.Default())

// 	v1 := r.Group("api")
// 	{
// 		nq := controller.NewQueueController()

// 		c := v1.Group("/Queue")
// 		{
// 			c.GET(":id", nq.ReadQueue)
// 			c.POST("", nq.CreateQueue)
// 			c.GET("", nq.ReadAllQueue)
// 			// c.PUT("", nc.UpdateCustomer)
// 			c.DELETE(":id", nq.DeleteQueue)
// 		}

// 	}
// 	r.Run(":8086")

// }

func main() {
	// repository.AutoMigrate()
	// repository.GenerateCode("A")
	// repository.CreateQueue("A")
	repository.GetAllQueues1()
	// repository.GetQueuesByType("C")
}
