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
	ReportQueue() (*model.ReportQueue, error)
	//line
	GetQueueLine(Code string) (*model.QueueResponseLine, error)
	DeleteQueuebyUID(UserID string) (*model.QueueResponse, error)
	FlexQueue(UserCode string) (string, error)
	AmountQueue(UserID string) (int, error)
	FlexReportQueue() (string, error)
}
