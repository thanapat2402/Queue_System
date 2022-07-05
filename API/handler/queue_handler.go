package handler

import (
	"net/http"
	"q/model"
	"q/service"

	"github.com/gin-gonic/gin"
)

type queueHandler struct {
	qService service.QueueService
}

func NewQueueHandler(qService service.QueueService) QueueHandler {
	return queueHandler{qService: qService}
}

func (h queueHandler) GetQueues(c *gin.Context) {
	queues, err := h.qService.GetQueues()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queues})
}

func (h queueHandler) GetQueuesType(c *gin.Context) {
	genre := c.Param("Type")
	if genre == "A" || genre == "B" || genre == "C" || genre == "D" {
		queues, err := h.qService.GetQueuesType(c.Param("Type"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": queues})
		return
	}
	c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid types error"})
}

func (h queueHandler) SearchQueue(c *gin.Context) {
	name := c.Query("name")
	genre := c.Query("types")
	if genre == "A" || genre == "B" || genre == "C" || genre == "D" {
		queues, err := h.qService.SearchQueue(name, genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": queues})
		return
	}
	c.JSON(http.StatusNotAcceptable, gin.H{"error": "invalid types error"})
}

func (h queueHandler) GetQueue(c *gin.Context) {
	queue, err := h.qService.GetQueue(c.Param("Code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queue})
}

func (h queueHandler) AddQueue(c *gin.Context) {
	var input model.QueueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	queue, err := h.qService.AddQueue(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queue, "message": "Created"})
}

func (h queueHandler) DeQueue(c *gin.Context) {
	queue, err := h.qService.DeQueue(c.Param("Code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queue, "message": "Deleted"})
}