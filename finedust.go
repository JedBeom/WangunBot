package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JedBeom/airq"
)

type HangulQ struct {
	Pm10 string
	Pm25 string
}

var (
	hangulQ HangulQ
)

func getAirq() {
	err := airq.LoadServiceKey("airq_key.txt")
	if err != nil {
		log.Println(err)
	}

	quality, err := airq.GetAirqOfNowByStation("연향동")
	if err != nil {
		log.Println(err)
	}

	var rate string
	switch quality.Pm10GradeWHO {
	case 1:
		rate = "최고"
	case 2:
		rate = "좋음"
	case 3:
		rate = "양호"
	case 4:
		rate = "보통"
	case 5:
		rate = "나쁨"
	case 6:
		rate = "상당히 나쁨"
	case 7:
		rate = "매우 나쁨"
	case 8:
		rate = "최악"
	}
	hangulQ.Pm10 = rate

	switch quality.Pm25GradeWHO {
	case 1:
		rate = "최고"
	case 2:
		rate = "좋음"
	case 3:
		rate = "양호"
	case 4:
		rate = "보통"
	case 5:
		rate = "나쁨"
	case 6:
		rate = "상당히 나쁨"
	case 7:
		rate = "매우 나쁨"
	case 8:
		rate = "최악"
	}
	hangulQ.Pm25 = rate

	return

}

func sendAirq(w http.ResponseWriter) {
	template := "현재 순천왕운중학교 주위 공기 상태는\n미세먼지는 %s, 초미세먼지는 %s!"
	content := fmt.Sprintf(template, hangulQ.Pm10, hangulQ.Pm25)

	message := Message{Text: content}
	keyboard := Keyboard{Type: "buttons", Buttons: buttons}
	response := Response{Message: message, Keyboard: keyboard}

	b, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Write(b)

	return

}
