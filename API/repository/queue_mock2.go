package repository

import (
	"errors"
	"q/model"
	"strconv"
	"strings"
	"time"
)

type queueRepositoryMock2 struct {
	queues []model.QueueModel
}

func NewQueueRepositoryMock2() QueueRepository {
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
	return queueRepositoryMock2{queues: queues}
}

func (r queueRepositoryMock2) GetAllQueues() ([]model.QueueModel, error) {
	return r.queues, nil
}

func (r queueRepositoryMock2) GetQueuesByType(types string) ([]model.QueueModel, error) {
	qOutput := []model.QueueModel{}
	for _, queues := range r.queues {
		if queues.Type == types {
			qOutput = append(qOutput, queues)
		}
	}
	return qOutput, nil
}

func (r queueRepositoryMock2) GetQueuesByNameTypes(name string, types string) (*model.QueueModel, error) {
	for _, queues := range r.queues {
		if queues.Name == name && queues.Type == types {
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r queueRepositoryMock2) GetQueuesByCode(strcode string) (*model.QueueModel, error) {
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	for _, queues := range r.queues {
		if queues.Code == code && queues.Type == Type {
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r queueRepositoryMock2) DeleteQueue(strcode string) (*model.QueueModel, error) {
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)

	for _, queues := range r.queues {
		if queues.Code == code && queues.Type == Type { //sep type
			return &queues, nil
		}
	}
	return nil, errors.New("record not found")
}

func (r queueRepositoryMock2) CreateQueue(data model.QueueInput) (*model.QueueModel, error) {
	date := time.Now()
	Queue := model.QueueModel{
		Code: 1,
		Type: data.Type,
		Date: date,
		Name: data.Name,
		Tel:  data.Tel}
	return &Queue, nil
}
