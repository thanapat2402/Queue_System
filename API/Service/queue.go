package service

import (
	"q/model"
)

//port
type QueueService interface {
	GetQueues() ([]model.QueueResponse, error)
	GetQueuesType(Type string) ([]model.QueueResponse, error)
	GetQueue(Code string) (*model.QueueResponse, error)
	AddQueue(data model.QueueInput) (*model.QueueResponse, error)
	DeQueue(Code string) (*model.QueueResponse, error)
}