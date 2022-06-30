package service

//Business logic

import (
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

func (s queueService) GetQueues() ([]model.QueueResponse, error) {
	queues, err := s.queueRepo.GetAllQueues()
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	qReponses := []model.QueueResponse{}
	for _, queue := range queues {
		qReponse := model.QueueResponse{
			Code: queue.Code,
			Date: queue.Date,
			Name: queue.Name,
			Tel:  queue.Tel,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s queueService) GetQueuesType(Type string) ([]model.QueueResponse, error) {
	queues, err := s.queueRepo.GetQueuesByType(Type)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	qReponses := []model.QueueResponse{}
	for _, queue := range queues {
		qReponse := model.QueueResponse{
			Code: queue.Code,
			Date: queue.Date,
			Name: queue.Name,
			Tel:  queue.Tel,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s queueService) GetQueue(code string) (*model.QueueResponse, error) {
	queue ,err := s.queueRepo.GetQueuesByCode(code)
	if err != nil{
		log.Println(err)
		return nil,err
	}
	qReponse := model.QueueResponse{
		Code: queue.Code,
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}

func (s queueService) AddQueue(data model.QueueInput) (*model.QueueResponse, error) {
	queue ,err := s.queueRepo.CreateQueue(data)
	if err != nil{
		log.Println(err)
		return nil,err
	}
	qReponse := model.QueueResponse{
		Code: queue.Code,
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}

func (s queueService) DeQueue(code string) (*model.QueueResponse, error) {
	queue ,err := s.queueRepo.DeleteQueue(code)
	if err != nil{
		log.Println(err)
		return nil,err
	}
	qReponse := model.QueueResponse{
		Code: queue.Code,
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}