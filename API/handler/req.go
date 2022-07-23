package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"q/model"
)

func Getprofile() (detail *model.LineProfile) {

	url := "https://api.line.me/v2/bot/profile/U75d559eb17b924479b63d01491314f48"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer 8AgOT7UHy3ypdFinyOdtM5/fR1rv85e+LRrKf0fsauW+SN6m5YmJ2tL5l9HEPAyrHStpGd7S2zoYYA16EdOm2n/rgf0HPYqz8R0xxSwIWqQkRbrq7buWdlalJLxzxlQQmiXABHjQYMm+tl0jJypb1AdB04t89/1O/w1cDnyilFU=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	data := model.LineProfile{}

	_ = json.Unmarshal([]byte(body), &data)

	return &data
}
