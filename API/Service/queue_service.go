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
		if err.Error() == "user Code not found" {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, ErrRepository
	} else {
		qReponse := model.QueueResponse{
			Code:   fmt.Sprintf("%v%03d", queue.Type, queue.Code),
			Date:   queue.Date,
			Name:   queue.Name,
			Tel:    queue.Tel,
			UserID: queue.UserID,
		}
		return &qReponse, nil
	}
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
		currentA = "no queue"
	} else {
		currentA = a[0].Code
	}
	if b == nil {
		currentB = "no queue"
	} else {
		currentB = b[0].Code
	}
	if c == nil {
		currentC = "no queue"
	} else {
		currentC = c[0].Code
	}
	if d == nil {
		currentD = "no queue"
	} else {
		currentD = d[0].Code
	}
	list := []int{len(a), len(b), len(c), len(d)}
	Amount := []string{}
	for _, i := range list {
		if i == 0 {
			Amount = append(Amount, "don't wait")
		} else if i == 1 {
			Amount = append(Amount, "a queue")
		} else {
			Amount = append(Amount, fmt.Sprintf("%v queues", i))
		}
	}
	qReport := model.ReportQueue{
		AmountQueueA:  Amount[0],
		AmountQueueB:  Amount[1],
		AmountQueueC:  Amount[2],
		AmountQueueD:  Amount[3],
		CurrentqueueA: currentA,
		CurrentqueueB: currentB,
		CurrentqueueC: currentC,
		CurrentqueueD: currentD,
	}

	return &qReport, nil
}
