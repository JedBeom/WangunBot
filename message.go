package main

import (
	"encoding/json"
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

	switch post.Content {

	case "홈":
		sendMessage(w, "저에게 무슨 요청을 하실 건가요?", home)
		return

	case "급식":
		sendMessage(w, "요일을 선택해주세요.", weekdays)
		return

	case "월요일", "화요일", "수요일", "목요일", "금요일":
		sendMeal(w, post)
		return

	case "미세먼지":
		sendAirq(w)
		return

	case "일정":
		dDay(w)
		return

	case "피드백":
		message := "왕운봇이 더 개선되기 위한 방안을 적어주세요!\n욕설 사용 시 법적조치 됩니다.\n뒤로 돌아가려면 '홈'을 입력해 주세요."

		sendMessage(w, message, []string{})

		return

	default:
		// 피드백 받았을 때의 경우겠죠 그런 경우 밖에 없음 아무튼 그럼

		// feedback.log에 저장
		feedback.Println(post.UserKey, post.Content)

		sendMessage(w, "피드백 감사합니다. 검토 후 반영을 결정 하겠습니다.", home)
		return
	}

}
