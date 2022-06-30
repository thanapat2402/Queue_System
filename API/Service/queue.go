package service

import (
	"q/model"
)

//port
type QueueService interface {
	GetQueues() ([]model.QueuesResponse, error)
	GetQueuesType(Type string) ([]model.QueuesResponse, error)
	GetQueue(Code string) (*model.QueueResponse, error)
	AddQueue(data model.QueueInput) (*model.QueueResponse, error)
	DeQueue(Code string) (*model.QueueResponse, error)
}