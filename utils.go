package main

import (
	"fmt"
	"net/http"
)

func sendMessage(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	template := `{
	"message":{
		"text": "%v"
	},
	"keyboard": {
		"type": "buttons",
		"buttons": ["월요일", "화요일", "수요일", "목요일", "금요일"]
	}
}`

	response := fmt.Sprintf(template, message)
	w.Write([]byte(response))
	return
}
