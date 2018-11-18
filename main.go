package main

import (
	"encoding/json"
	"net/http"
	"time"

	sm "github.com/JedBeom/schoolmeal"
)

func getMealByWeekday(weekday time.Weekday) (meal string, err error) {
	school := sm.School{
		SchoolCode:     "Q100005451",
		SchoolKindCode: sm.Middle,
		Zone:           sm.Jeonnam,
	}

	meals, err := school.GetWeekMeal(sm.Timestamp(), sm.Lunch)
	if err != nil {
		return
	}

	meal = meals[weekday]

	return
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var post Post

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var meal string
	switch post.Content {
	case "월요일":
		meal, err = getMealByWeekday(time.Monday)
	case "화요일":
		meal, err = getMealByWeekday(time.Tuesday)
	case "수요일":
		meal, err = getMealByWeekday(time.Wednesday)
	case "목요일":
		meal, err = getMealByWeekday(time.Thursday)
	case "금요일":
		meal, err = getMealByWeekday(time.Friday)
	}

	if meal == " " {
		meal = "그 날의 급식이 없어요!"
	}

	if err != nil {
		meal = "왕운봇에 문제가 있는 것 같네요. 금방 고쳐질 거에요!"
	}

	sendMessage(w, meal)
}

func keyboardHandler(w http.ResponseWriter, r *http.Request) {
	keyboard := `{
	"type": "buttons",
	"buttons": ["월요일", "화요일", "수요일", "목요일", "금요일"]
}`

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(keyboard))
}

func main() {
	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/keyboard", keyboardHandler)

	server.ListenAndServe()
}
