package service

import (
	"errors"
	"fmt"
	"log"
	"q/model"
)

func (s queueService) GetQueueLine(code string) (*model.QueueResponseLine, error) {
	queue, err := s.queueRepo.GetQueuesByCode(code)
	if err != nil {
		log.Println(err)
		return nil, errors.New("repository error")
	}
	current, err := s.queueRepo.GetCurrentQueue(queue.Type)
	if err != nil {
		log.Println(err)
		return nil, errors.New("repository error")
	}
	qReponse := model.QueueResponseLine{
		CurrentCode: fmt.Sprintf("%v%03d", current.Type, current.Code),
		UserCode:    fmt.Sprintf("%v%03d", queue.Type, queue.Code),
		QueueAmount: queue.Code - current.Code,
		Date:        queue.Date,
		Name:        queue.Name,
	}
	fmt.Println(qReponse)
	return &qReponse, nil
}

func (s queueService) DeleteQueuebyUID(UserID string) (*model.QueueResponse, error) {
	queue, err := s.queueRepo.DeleteQueuebyUID(UserID)
	if err != nil {
		if err.Error() == "user Code not found" {
			return nil, err
		}
		log.Println(err)
		return nil, errors.New("repository error")
	} else {
		qReponse := model.QueueResponse{
			Code:   fmt.Sprintf("%v%03d", queue.Type, queue.Code),
			Date:   queue.Date,
			Name:   queue.Name,
			Tel:    queue.Tel,
			UserID: queue.UserID,
		}
		log.Printf("%v is cancle queue", queue.Name)
		return &qReponse, nil
	}
}

func (s queueService) AmountQueue(UserID string) (int, error) {

	User_queue, err := s.queueRepo.GetQueueByUserID(UserID)
	if err != nil {
		log.Println(err)
		return 99999, errors.New("repository error")
	}
	queues, err := s.queueRepo.GetQueuesByType(User_queue.Type)
	if err != nil {
		log.Println(err)
		return 99999, errors.New("repository error")
	}
	wait := 0
	for _, queue := range queues {
		if queue.Code < User_queue.Code {
			wait += 1
		}
	}
	log.Printf("%v is check queue", User_queue.Name)
	return wait, nil
}

func (s queueService) FlexQueue(UserCode string) (string, error) {
	queue, err := s.GetQueueLine(UserCode)
	if err != nil {
		return "", err
	}
	var wait string
	if queue.QueueAmount == 1 {
		wait = "Waiting a queue"
	} else if queue.QueueAmount > 1 {
		wait = fmt.Sprintf("Waiting %v queues", queue.QueueAmount)
	} else if queue.QueueAmount == 0 {
		wait = "It's your turn"
	}
	if queue.Name == "" {
		queue.Name = "ไม่ระบุชื่อ"
	}
	flex := fmt.Sprintf(QueueFlex, queue.UserCode, queue.Name, queue.Date.Format("Monday 2, 15:04:05"), wait, queue.UserCode)

	return flex, nil
}

func (s queueService) FlexReportQueue() (string, error) {
	report, err := s.ReportQueue()
	if err != nil {
		return "", err
	}
	flex := fmt.Sprintf(ReportFlex, report.CurrentqueueA, report.AmountQueueA, report.CurrentqueueB, report.AmountQueueB, report.CurrentqueueC, report.AmountQueueC, report.CurrentqueueD, report.AmountQueueD)
	return flex, nil
}

// func pushmessage (userID string,message string){
// 	bot, err := linebot.New(<channel secret>, <channel token>)
// 	if err != nil {
// 	...
// 	}
// 	if _, err := bot.Multicast(userIDs, linebot.NewTextMessage("hello")).Do(); err != nil {
// 	...
// 	}
// }
