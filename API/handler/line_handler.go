package handler

import (
	"fmt"
	"log"
	"net/http"
	"q/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func (h queueHandler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (h queueHandler) Callback(c *gin.Context) {
	bot := GetBot()
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypePostback {
			data := event.Postback.Data
			rid1, rid2 := ReadMenu()
			fmt.Println(data)
			switch data {
			case "RichMenu1":
				if _, err = bot.LinkUserRichMenu(event.Source.UserID, rid1).Do(); err != nil {
					log.Fatal(err)
				}
			case "RichMenu2":
				if _, err = bot.LinkUserRichMenu(event.Source.UserID, rid2).Do(); err != nil {
					log.Fatal(err)
				}
			}
			return
		}
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "test" {

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("test")).Do(); err != nil {
						log.Print(err)
					}
				}
				if message.Text == "ยกเลิกคิว" {
					queue, err := h.qService.DeleteQueuebyUID(event.Source.UserID)
					if err != nil {
						if err.Error() == "user Code not found" {
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("ท่านยังไม่ได้จองคิวไม่สามารถยกเลิกได้")).Do(); err != nil {
								log.Print(err)
							}
							return
						} else {
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("เกิดข้อผิดพลาดไม่สามารถยกเลิกคิวได้")).Do(); err != nil {
								log.Print(err)
							}
							return
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("ท่านยกเลิกคิว %v เรียบร้อยแล้ว", queue.Code))).Do(); err != nil {
						log.Print(err)
					}
					return
				}

				if message.Text == "จองคิว" {
					// Unmarshal JSON
					flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(MenuFlex))
					if err != nil {
						log.Println(err)
					}
					fmt.Println(flexContainer)
					// New Flex Message
					flexMessage := linebot.NewFlexMessage(message.Text, flexContainer)
					// Reply Message
					_, err = bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
					if err != nil {
						log.Print(err)
					}
					return
				}
				if message.Text == "Alone" || message.Text == "Couple" || message.Text == "Small Group" || message.Text == "The Gang" || message.Text == "VVIP" {
					var types string
					switch message.Text {
					case "Alone":
						types = "A"
					case "Couple":
						types = "B"
					case "Small Group":
						types = "C"
					case "The Gang":
						types = "D"
					case "VVIP":
						types = "V"
					default:
						log.Println("This Type not in Conditions")
					}

					input := model.QueueInput{
						Type:   types,
						Name:   Getprofile(event.Source.UserID).DisplayName,
						UserID: event.Source.UserID,
					}
					queue, err := h.qService.AddQueue(input)
					if err != nil {
						if err.Error() == "queue already exists" {
							if _, err := bot.PushMessage(input.UserID, linebot.NewTextMessage("ท่านจองคิวไปแล้วกรุณายกเลิกคิวก่อนหน้า")).Do(); err != nil {
								log.Print(err)
							}
							return
						} else {
							if _, err := bot.PushMessage(input.UserID, linebot.NewTextMessage("เกิดข้อผิดพลาดไม่สามารถบันทึกคิวได้")).Do(); err != nil {
								log.Print(err)
							}
							return
						}
					}
					if input.UserID != "" {
						flex, err := h.qService.FlexQueue(queue.Code)
						if err != nil {
							c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
							return
						}
						// Unmarshal JSON
						flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(flex))
						if err != nil {
							c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
							return
						}
						// New Flex Message
						flexMessage := linebot.NewFlexMessage("Your Queue", flexContainer)
						if _, err := bot.PushMessage(input.UserID, flexMessage).Do(); err != nil {
							c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
							return
						}
					}
					return
				}

				if message.Text == "ตรวจสอบคิวทั้งหมด" {
					ReportFlex, err := h.qService.FlexReportQueue()
					if err != nil {
						log.Println(err)
					}
					// Unmarshal JSON
					flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(ReportFlex))
					if err != nil {
						log.Println(err)
					}
					fmt.Println(flexContainer)
					// New Flex Message
					flexMessage := linebot.NewFlexMessage(message.Text, flexContainer)
					// Reply Message
					_, err = bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
					if err != nil {
						log.Print(err)
					}
					return
				}

				if message.Text == "ตรวจสอบคิวของฉัน" {
					wait, err := h.qService.AmountQueue(event.Source.UserID)
					if err != nil {
						log.Println(err)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่สามารถคำนวณจำนวณที่ต้องรอคิว คุณอาจยังไม่ได้จองคิว")).Do(); err != nil {
							log.Print(err)
						}
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(wait)).Do(); err != nil {
						log.Print(err)
					}
					return
				}

				if message.Text == "golf" {
					userIDs := "U75d559eb17b924479b63d01491314f48"
					if _, err := bot.PushMessage(userIDs, linebot.NewTextMessage("มีคนอยากเซ็ทหย่อสูดต่อซูดผ่อซีหม่อสองห่อใส่ไข่กับคุณ")).Do(); err != nil {
						log.Print(err)
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ส่งข้อความให้กอล์ฟแล้ว")).Do(); err != nil {
						log.Print(err)
						return
					}
				}

				split := strings.Split(message.Text, " ")
				if split[0] == "ดู" || split[0] == "ตรวจสอบ" || split[0] == "ค้นหา" {
					fmt.Println(split[1])
					flex, err := h.qService.FlexQueue(split[1])
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
					flexMessage := linebot.NewFlexMessage(split[1], flexContainer)
					// Reply Message
					_, err = bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
					if err != nil {
						log.Print(err)
					}
				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ขออภัยครับ แต่เรายังไม่เข้าใจ ท่านอยากจะทวนอีกรอบหรือส่งต่อให้เจ้าหน้าที่ตอบคำถามดีครับ")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
