package repository

import (
	"errors"
	"q/model"
	"strconv"
	"strings"
	"time"
)

type queueRepositoryMock struct {
	queues []model.QueueModel
}

func NewQueueRepositoryMock() QueueRepository {
	queues := []model.QueueModel{
		{
			Code: 4,
			Type: "A",
			Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			Name: "Nop",
			Tel:  "1112"},
		{
			Code: 5,
			Type: "B",
			Date: time.Date(2022, time.July, 12, 21, 34, 05, 0, time.UTC),
			Name: "Steven",
			Tel:  "191"},
	}
	return queueRepositoryMock{queues: queues}
}

func (r queueRepositoryMock) GetAllQueues() ([]model.QueueModel, error) {
	return r.queues, nil
}

func (r queueRepositoryMock) GetQueuesByType(types string) ([]model.QueueModel, error) {
	qOutput := []model.QueueModel{}
	for _, queues := range r.queues {
		if queues.Type == types {
			qOutput = append(qOutput, queues)
		}
	}
	return qOutput, nil
}

func (r queueRepositoryMock) GetQueuesByCode(strcode string) (*model.QueueModel, error) {
	code, _ := strconv.Atoi(strings.TrimLeft(strcode, "ABCD"))
	for _, queues := range r.queues {
		if queues.Code == code {
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

// func (r queueRepositoryMock) DeleteQueue(strcode string) (*model.QueueModel, error) {
// 	return nil, nil
// }

func (r queueRepositoryMock) DeleteQueue(strcode string) (*model.QueueModel, error) {
	code, _ := strconv.Atoi(strings.TrimLeft(strcode, "ABCD"))
	for _, queues := range r.queues {
		if queues.Code == code {
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r queueRepositoryMock) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	date := time.Now()
	Queue := model.QueueModel{
		Code: 1,
		Type: data.Type,
		Date: date,
		Name: data.Name,
		Tel:  data.Tel}
	return &Queue, nil
}
