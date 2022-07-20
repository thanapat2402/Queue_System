package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func (h queueHandler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (h queueHandler) Callback(c *gin.Context) {
	bot := GetBot()
	events, err := bot.ParseRequest(c.Request)
	fmt.Println(err)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		fmt.Println(event.Source.UserID)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				userIDs := "U75d559eb17b924479b63d01491314f48"
				if message.Text == "ยกเลิกคิว" {
					h.qService.DeleteQueuebyUID(event.Source.UserID)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ยกเลิกคิวเรียบร้อยแล้วครับ")).Do(); err != nil {
						log.Print(err)
					}
					return
				}
				if message.Text == "ตรวจสอบคิว" {
					wait, err := h.qService.AmountQueue(event.Source.UserID)
					if err != nil {
						log.Println(err)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่สามารถคำนวณจำนวณที่ต้องรอคิว คุณอาจยังไม่ได้จองคิว")).Do(); err != nil {
							log.Print(err)
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("กรุณารออีก %v คิว", wait))).Do(); err != nil {
						log.Print(err)
					}
					return
				}
				if message.Text == "golf" {
					if _, err := bot.PushMessage(userIDs, linebot.NewTextMessage("มีคนอยากเซ็ทหย่อสูดต่อซูดผ่อซีหม่อสองห่อใส่ไข่กับคุณ")).Do(); err != nil {
						log.Print(err)
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ส่งข้อความให้กอล์ฟแล้ว")).Do(); err != nil {
						log.Print(err)
						return
					}
				}
				flex, err := h.qService.FlexQueue(message.Text)
				fmt.Println(err)
				if err != nil {
					if err.Error() == "repository error" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่พบเลขคิวที่คุณค้นหาหรืออาจเลยคิวของคุณมาแล้ว")).Do(); err != nil {
							log.Print(err)
							return
						}
						return
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ระบบผิดพลาด")).Do(); err != nil {
							log.Print(err)
							return
						}
						return
					}
				}
				// Unmarshal JSON
				flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(flex))
				if err != nil {
					log.Println(err)
				}
				// New Flex Message
				flexMessage := linebot.NewFlexMessage("message.Text", flexContainer)
				// Reply Message
				_, err = bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
				if err != nil {
					log.Print(err)
				}

				// } else {
				// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				// 		log.Print(err)
				// 	}
				// }
			}
		}
	}
}
