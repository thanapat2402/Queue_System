package service

import (
	"q/model"
)

//port
type QueueService interface {
	//web
	GetQueues() ([]model.QueuesResponse, error)
	GetQueuesType(Type string) ([]model.QueuesResponse, error)
	SearchQueue(name string, types string) (*model.QueueResponse, error)
	GetQueue(Code string) (*model.QueueResponse, error)
	AddQueue(data model.QueueInput) (*model.QueueResponse, error)
	DeQueue(Code string) (*model.QueueResponse, error)
	//line
	GetQueueLine(Code string) (*model.QueueResponseLine, error)
	DeleteQueuebyUID(UserID string) error
	FlexQueue(UserID string) (string, error)
	AmountQueue(UserID string) (int, error)
}
