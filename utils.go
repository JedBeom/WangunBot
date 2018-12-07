package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 받은것만 보냅니다
func sendMessage(w http.ResponseWriter, content string, buttons []string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	keyboard := Keyboard{
		Type:    "buttons",
		Buttons: buttons,
	}

	if len(buttons) == 0 {
		keyboard.Type = "Text"
	}

	message := Message{
		Text: content,
	}

	response := Response{
		Keyboard: keyboard,
		Message:  message,
	}

	b, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Write(b)
	return
}

func logger(p Post) {
	log.Println(p.UserKey, p.Content)
}
