package service

import (
	"q/model"

	"github.com/stretchr/testify/mock"
)

type queueServiceMock struct {
	mock.Mock
}

func NewQueueServiceMock() *queueServiceMock {
	return &queueServiceMock{}
}

func (s *queueServiceMock) GetQueues() ([]model.QueuesResponse, error) {
	args := s.Called()
	return args.Get(0).([]model.QueuesResponse), args.Error(1)
}

func (r *queueServiceMock) GetQueuesType(types string) ([]model.QueuesResponse, error) {
	args := r.Called(types)
	return args.Get(0).([]model.QueuesResponse), args.Error(1)
}

func (r *queueServiceMock) SearchQueue(name string, types string) (*model.QueueResponse, error) {
	args := r.Called(name, types)
	return args.Get(0).(*model.QueueResponse), args.Error(1)
}

func (r *queueServiceMock) GetQueue(strcode string) (*model.QueueResponse, error) {
	args := r.Called(strcode)
	return args.Get(0).(*model.QueueResponse), args.Error(1)
}

func (r *queueServiceMock) AddQueue(data model.QueueInput) (*model.QueueResponse, error) {
	args := r.Called(data)
	return args.Get(0).(*model.QueueResponse), args.Error(1)
}

func (r *queueServiceMock) DeQueue(strcode string) (*model.QueueResponse, error) {
	args := r.Called(strcode)
	return args.Get(0).(*model.QueueResponse), args.Error(1)
}

func (r *queueServiceMock) GetQueueLine(code string) (*model.QueueResponseLine, error) {
	args := r.Called(code)
	return args.Get(0).(*model.QueueResponseLine), args.Error(1)
}

func (r *queueServiceMock) DeleteQueuebyUID(UserID string) error {
	args := r.Called(UserID)
	return args.Error(0)
}

func (r *queueServiceMock) FlexQueue(UserID string) (string, error) {
	args := r.Called(UserID)
	return args.Get(0).(string), args.Error(1)
}

func (r *queueServiceMock) CheckQueue(UserID string) (string, error) {
	args := r.Called(UserID)
	return args.Get(0).(string), args.Error(1)
}

func (r *queueServiceMock) AmountQueue(UserID string) (int, error) {
	args := r.Called(UserID)
	return args.Get(0).(int), args.Error(1)
}

func (s *queueServiceMock) ReportQueue() (*model.ReportQueue, error) {
	args := s.Called()
	return args.Get(0).(*model.ReportQueue), args.Error(1)
}
