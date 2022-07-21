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

func (s queueService) DeleteQueuebyUID(UserID string) error {
	queue, err := s.queueRepo.DeleteQueuebyUID(UserID)
	if err != nil {
		log.Println(err)
		return errors.New("repository error")
	}
	log.Printf("%v is cancle queue", queue.Name)
	return nil
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
	flex := fmt.Sprintf(`{
		"type": "bubble",
		"size": "kilo",
		"direction": "ltr",
		"hero": {
		  "type": "image",
		  "url": "https://www.i-pic.info/i/KMdp196143.png",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "position": "relative"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "md",
		  "contents": [
			{
			  "type": "text",
			  "text": "%v",
			  "weight": "bold",
			  "size": "xl",
			  "gravity": "center",
			  "margin": "lg",
			  "wrap": true,
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "Name",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "Date",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "margin": "xs",
			  "contents": [
				{
				  "type": "text",
				  "text": "Queue",
				  "size": "sm",
				  "color": "#AAAAAA",
				  "flex": 2,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%v",
				  "size": "sm",
				  "color": "#666666",
				  "flex": 4,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "spacer",
				  "size": "xs"
				},
				{
				  "type": "image",
				  "url": "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=%s",
				  "size": "md",
				  "aspectMode": "cover"
				},
				{
				  "type": "text",
				  "text": "You can enter the restaurant by using this code instead of a ticket",
				  "size": "xxs",
				  "color": "#AAAAAA",
				  "margin": "xxl",
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		}
	  }`, queue.UserCode, queue.Name, queue.Date.Format("Monday 2, 15:04:05"), wait, queue.UserCode)
	
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
