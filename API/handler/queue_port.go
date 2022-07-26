package handler

import "github.com/gin-gonic/gin"

type QueueHandler interface {
	//web
	GetQueues(c *gin.Context)
	GetQueuesType(c *gin.Context)
	SearchQueue(c *gin.Context)
	GetQueue(c *gin.Context)
	AddQueue(c *gin.Context)
	DeQueue(c *gin.Context)
	DeQueue2(c *gin.Context)
	ReportQueue(c *gin.Context)
	//line
	Callback(c *gin.Context)
	Hello(c *gin.Context)
}
