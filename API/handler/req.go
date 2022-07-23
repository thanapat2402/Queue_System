package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"q/model"
)

func Getprofile(userID string) (detail *model.LineProfile) {
	url := fmt.Sprintf("https://api.line.me/v2/bot/profile/%v", userID)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, token := Readline()
	auth := fmt.Sprintf("Bearer %v", token)
	req.Header.Add("Authorization", auth)

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
