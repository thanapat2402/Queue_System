package repository

import (
	"errors"
	"q/model"
	"strconv"
	"strings"
	"time"

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
	print(args)
	return args.Get(0).([]model.QueueModel), args.Error(1)
}

func (r *queueRepositoryMock) GetQueuesByType(types string) ([]model.QueueModel, error) {
	args := r.Called(types)
	qOutput := []model.QueueModel{}
	queues := args.Get(0).([]model.QueueModel)
	for _, queue := range queues {
		if queue.Type == types {
			qOutput = append(qOutput, queue)
		}
	}
	return qOutput, nil
}

func (r *queueRepositoryMock) SearchQueuesByNameTypes(name string, types string) (*model.QueueModel, error) {
	args := r.Called(name, types)
	queues := args.Get(0).([]model.QueueModel)
	for _, queues := range queues {
		if queues.Name == name && queues.Type == types {
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r *queueRepositoryMock) GetQueuesByCode(strcode string) (*model.QueueModel, error) {
	args := r.Called(strcode)
	queues := args.Get(0).([]model.QueueModel)

	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	for _, queues := range queues {
		if queues.Code == code && queues.Type == Type {
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r *queueRepositoryMock) DeleteQueue(strcode string) (*model.QueueModel, error) {
	args := r.Called(strcode)
	queues := args.Get(0).([]model.QueueModel)

	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)

	for _, queues := range queues {
		if queues.Code == code && queues.Type == Type { //sep type
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r *queueRepositoryMock) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	date := time.Now()
	Queue := model.QueueModel{
		Code: 1,
		Type: data.Type,
		Date: date,
		Name: data.Name,
		Tel:  data.Tel}
	return &Queue, nil
}
