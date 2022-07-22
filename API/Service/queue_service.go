package service

//Business logic in here

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
		log.Panic(ErrRepository)
		log.Println(err)
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
		log.Println(err)
		return nil, ErrRepository
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

func (s queueService) SearchQueue(name string, types string) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.GetQueuesByNameTypes(name, types)
	if err != nil {
		log.Println(err)
		return nil, ErrRepository
	}
	qReponse := model.QueueResponse{
		Code: fmt.Sprintf("%v%03d", queue.Type, queue.Code),
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}

func (s queueService) GetQueue(code string) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.GetQueuesByCode(code)
	if err != nil {
		log.Println(err)
		return nil, ErrRepository
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
		if err.Error() == "queue already exists" {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, ErrRepository
	} else {
		code := fmt.Sprintf("%v%03d", queue.Type, queue.Code)
		qReponse := model.QueueResponse{
			Code: code,
			Date: queue.Date,
			Name: queue.Name,
			Tel:  queue.Tel,
		}
		return &qReponse, nil
	}
}

func (s queueService) DeQueue(code string) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.DeleteQueue(code)
	if err != nil {
		log.Println(err)
		return nil, ErrRepository
	}
	qReponse := model.QueueResponse{
		Code: fmt.Sprintf("%v%03d", queue.Type, queue.Code),
		Date: queue.Date,
		Name: queue.Name,
		Tel:  queue.Tel,
	}
	return &qReponse, nil
}

func (s queueService) ReportQueue() (*model.ReportQueue, error) {
	queues, err := s.GetQueues()
	println(err)
	if err != nil {
		log.Panic(ErrRepository)
		return nil, err
	}
	var a, b, c, d []model.QueuesResponse
	for _, queue := range queues {
		switch queue.Type {
		case "A":
			a = append(a, queue)
		case "B":
			b = append(b, queue)
		case "C":
			c = append(c, queue)
		case "D":
			d = append(d, queue)
		default:
			log.Println("This Type not in Conditions")
		}
	}

	var currentA, currentB, currentC, currentD string
	if a == nil {
		currentA = "no queue type A"
	} else {
		currentA = a[0].Code
	}
	if b == nil {
		currentB = "no queue type B"
	} else {
		currentB = b[0].Code
	}
	if c == nil {
		currentC = "no queue type C"
	} else {
		currentC = c[0].Code
	}
	if d == nil {
		currentD = "no queue type D"
	} else {
		currentD = d[0].Code
	}
	qReport := model.ReportQueue{
		AmountQueueA:  len(a),
		AmountQueueB:  len(b),
		AmountQueueC:  len(c),
		AmountQueueD:  len(d),
		CurrentqueueA: currentA,
		CurrentqueueB: currentB,
		CurrentqueueC: currentC,
		CurrentqueueD: currentD,
	}

	return &qReport, nil
}
