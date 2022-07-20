package repository

import (
	"q/model"
)

//Port
type QueueRepository interface {
	//web
	GetAllQueues() ([]model.QueueModel, error)
	GetQueuesByType(Type string) ([]model.QueueModel, error)
	GetQueuesByNameTypes(name string, types string) (*model.QueueModel, error)
	GetQueuesByCode(Code string) (*model.QueueModel, error)
	CreateQueue(data model.QueueInput) (*model.QueueModel, error)
	DeleteQueue(Code string) (*model.QueueModel, error)
	//line
	GetCurrentQueue(types string) (*model.QueueModel, error)
	GetQueueByUserID(UserID string) (*model.QueueModel, error)
	DeleteQueuebyUID(UserID string) (*model.QueueModel, error)
}
