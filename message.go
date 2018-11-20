package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var post Post

	// 해독
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	logger(post)

	// 급식일 경우에...
	var meal string

	switch post.Content {

	case "급식":
		keyboard := Keyboard{Type: "buttons", Buttons: weekdays}
		message := Message{Text: "요일을 선택 해주세요."}
		response := Response{Keyboard: keyboard, Message: message}
		b, err := json.MarshalIndent(response, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Write(b)
		return

	case "미세먼지":
		sendAirq(w)
		return

	case "월요일":
		meal = meals[1]
	case "화요일":
		meal = meals[2]
	case "수요일":
		meal = meals[3]
	case "목요일":
		meal = meals[4]
	case "금요일":
		meal = meals[5]

	case "피드백":
		keyboard := Keyboard{
			Type:    "text",
			Buttons: []string{},
		}

		message := Message{Text: "왕운봇이 더 개선되기 위한 방안을 적어주세요!\n욕설 사용 시 법적조치 됩니다."}
		response := Response{Keyboard: keyboard, Message: message}

		b, err := json.MarshalIndent(response, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Write(b)

		return

	default:
		// 피드백 받았을 때의 경우겠죠 그런 경우 밖에 없음 아무튼 그럼

		// feedback.log에 저장
		feedback.Println(post.UserKey, post.Content)

		sendMessage(w, "피드백 감사합니다. 검토 후 반영을 결정 하겠습니다.")
		return
	}

	// 슬라이스에 저장된게 없을 때!
	if meal == " " || meal == "" {
		meal = "그 날의 급식이 없어요!"
	}

	sendMessage(w, meal)
}
