package handler

import "github.com/gin-gonic/gin"

type QueueHandler interface {
	GetQueues(c *gin.Context)
	GetQueuesType(c *gin.Context)
	SearchQueue(c *gin.Context)
	GetQueue(c *gin.Context)
	AddQueue(c *gin.Context)
	DeQueue(c *gin.Context)
}
