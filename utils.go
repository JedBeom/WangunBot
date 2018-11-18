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
	}
}`

	response := fmt.Sprintf(template, message)
	w.Write([]byte(response))
	return
}
