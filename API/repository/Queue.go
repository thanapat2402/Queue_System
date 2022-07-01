package repository

import (
	"q/model"
)

//Port
type QueueRepository interface {
	GetAllQueues() ([]model.QueueModel, error)
	GetQueuesByType(Type string) ([]model.QueueModel, error)
	SearchQueuesByNameTypes(name string, types string) (*model.QueueModel, error)
	GetQueuesByCode(Code string) (*model.QueueModel, error)
	CreateQueue(data model.QueueInput) (*model.QueueModel, error)
	DeleteQueue(Code string) (*model.QueueModel, error)
}
