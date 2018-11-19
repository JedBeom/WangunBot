package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	sm "github.com/JedBeom/schoolmeal"
	"github.com/jasonlvhit/gocron"
)

var (
	school = sm.School{
		SchoolCode:     "Q100005451",
		SchoolKindCode: sm.Middle,
		Zone:           sm.Jeonnam,
	}

	meals    []string
	feedback *log.Logger
)

func getMeals() {

	todayMeals, err := school.GetWeekMeal(sm.Timestamp(), sm.Lunch)
	if err != nil {
		log.Println(err)
		return
	}

	meals = todayMeals

}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var post Post

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	logger(post)

	var meal string
	switch post.Content {
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
		feedback.Println(post.UserKey, post.Content)

		sendMessage(w, "피드백 감사합니다. 검토 후 반영을 결정 하겠습니다.")
		return
	}

	if meal == " " || meal == "" {
		meal = "그 날의 급식이 없어요!"
	}

	sendMessage(w, meal)
}

func keyboardHandler(w http.ResponseWriter, r *http.Request) {
	keyboard := Keyboard{
		Type:    "buttons",
		Buttons: buttons,
	}

	b, err := json.MarshalIndent(keyboard, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func main() {
	port := ":" + os.Args[1]
	server := http.Server{
		Addr: port,
	}

	inputLog, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	defer inputLog.Close()

	log.Println("Starting")
	log.SetOutput(inputLog)

	feedbackLog, err := os.OpenFile("feedback.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	feedback = log.New(feedbackLog, "", 0)

	gocron.Every(1).Day().At("00:00").Do(getMeals)

	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/keyboard", keyboardHandler)
	getMeals()

	go server.ListenAndServe()
	<-gocron.Start()
}
