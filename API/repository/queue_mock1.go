package repository

import (
	"q/model"

	"github.com/stretchr/testify/mock"
)

type queueRepositoryMock struct {
	mock.Mock
}

func NewQueueRepositoryMock() *queueRepositoryMock {
	return &queueRepositoryMock{}
}

func (r *queueRepositoryMock) GetAllQueues() ([]model.QueueModel, error) {
	args := r.Called()
	return args.Get(0).([]model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) GetQueuesByType(types string) ([]model.QueueModel, error) {
	args := r.Called(types)
	return args.Get(0).([]model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) GetQueuesByNameTypes(name string, types string) (*model.QueueModel, error) {
	args := r.Called(name, types)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) GetQueuesByCode(strcode string) (*model.QueueModel, error) {
	args := r.Called(strcode)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) DeleteQueue(strcode string) (*model.QueueModel, error) {
	args := r.Called(strcode)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	args := r.Called(data)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) GetCurrentQueue(types string) (*model.QueueModel, error) {
	args := r.Called(types)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) DeleteQueuebyUID(UserID string) (*model.QueueModel, error) {
	args := r.Called(UserID)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) GetQueueByUserID(UserID string) (*model.QueueModel, error) {
	args := r.Called(UserID)
	return args.Get(0).(*model.QueueModel), args.Error(1)
}
