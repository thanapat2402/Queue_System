package service_test

import (
	"q/model"
	"q/repository"
	"q/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetQueues(t *testing.T) {
	//Arrage
	queueRepoMock := repository.NewQueueRepositoryMock2()
	queueService := service.NewQueueService(queueRepoMock)

	//Act
	queues, err := queueService.GetQueues()
	if err != nil {
		println(err)
	}
	expected := []model.QueuesResponse{
		{
			Code: "A004",
			Type: "A",
			Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			Name: "Nop",
			Tel:  "1112"},
		{
			Code: "B005",
			Type: "B",
			Date: time.Date(2022, time.July, 12, 21, 34, 05, 0, time.UTC),
			Name: "Steven",
			Tel:  "191"},
	}

	//Assert
	assert.Equal(t, expected, queues)

}

func TestGetAllQueues(t *testing.T) {
	//Arrange
	queueRepo := repository.NewQueueRepositoryMock()
	result := []model.QueueModel{
		{
			Code: 3,
			Type: "A",
			Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC),
			Name: "Golf",
			Tel:  "1150"},
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

	queueRepo.On("GetAllQueues").Return(result, nil)

	queueService := service.NewQueueService(queueRepo)

	//Act
	read, _ := queueService.GetQueues()
	println(read)
	expected := []model.QueuesResponse{
		{
			Code: "A003",
			Type: "A",
			Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC),
			Name: "Golf",
			Tel:  "1150"},
		{
			Code: "A004",
			Type: "A",
			Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			Name: "Nop",
			Tel:  "1112"},
		{
			Code: "B005",
			Type: "B",
			Date: time.Date(2022, time.July, 12, 21, 34, 05, 0, time.UTC),
			Name: "Steven",
			Tel:  "191"},
	}

	//Assert
	assert.Equal(t, expected, read)

}

func TestGetByType(t *testing.T) {
	//Arrange
	queueRepo := repository.NewQueueRepositoryMock()
	types := "A"
	queueRepo.On("GetQueuesByType", types).Return([]model.QueueModel{
		{
			Code: 3,
			Type: "A",
			Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC),
			Name: "Golf",
			Tel:  "1150"},
		{
			Code: 4,
			Type: "A",
			Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			Name: "Nop",
			Tel:  "1112"},
	}, nil)

	queueService := service.NewQueueService(queueRepo)
	//Act
	read, _ := queueService.GetQueuesType(types)
	expected := []model.QueuesResponse{
		{
			Code: "A003",
			Type: "A",
			Date: time.Date(2020, time.April, 10, 21, 34, 01, 0, time.UTC),
			Name: "Golf",
			Tel:  "1150"},
		{
			Code: "A004",
			Type: "A",
			Date: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			Name: "Nop",
			Tel:  "1112"},
	}

	//Assert
	assert.Equal(t, expected, read)

}