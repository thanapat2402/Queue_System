package service

//Business logic

import (
	"fmt"
	"log"
	"q/model"
	"q/repository"
)

type queueService struct {
	queueRepo repository.QueueRepository //อ้างถึง interface
}

//constructor
func NewQueueService(queueRepo repository.QueueRepository) QueueService {
	return queueService{queueRepo: queueRepo}
}

func (s queueService) GetQueues() ([]model.QueuesResponse, error) {
	queues, err := s.queueRepo.GetAllQueues()
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	qReponses := []model.QueuesResponse{}
	for _, queue := range queues {
		qReponse := model.QueuesResponse{
			Code: fmt.Sprintf("%v%03d", queue.Type, queue.Code),
			Date: queue.Date,
			Type: queue.Type,
			Name: queue.Name,
			Tel:  queue.Tel,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s queueService) GetQueuesType(Type string) ([]model.QueuesResponse, error) {
	queues, err := s.queueRepo.GetQueuesByType(Type)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	qReponses := []model.QueuesResponse{}
	for _, queue := range queues {
		qReponse := model.QueuesResponse{
			Code: fmt.Sprintf("%v%03d", queue.Type, queue.Code),
			Date: queue.Date,
			Type: queue.Type,
			Name: queue.Name,
			Tel:  queue.Tel,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s queueService) GetQueue(code string) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.GetQueuesByCode(code)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	qReponse := model.QueueResponse{
		Code: fmt.Sprintf("%v%03d", queue.Type, queue.Code),
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}

func (s queueService) AddQueue(data model.QueueInput) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.CreateQueue(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	code := fmt.Sprintf("%v%03d", queue.Type, queue.Code)
	qReponse := model.QueueResponse{
		Code: code,
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}

func (s queueService) DeQueue(code string) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.DeleteQueue(code)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	qReponse := model.QueueResponse{
		Code: fmt.Sprintf("%v%03d", queue.Type, queue.Code),
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}
