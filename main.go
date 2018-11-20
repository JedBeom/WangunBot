package main

import (
	"encoding/json"
	"fmt"
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

	// 급식 저장용
	meals []string

	// feedback.log
	feedback *log.Logger
)

// 급식을 불러옴
func getMeals() {

	todayMeals, err := school.GetWeekMeal(sm.Timestamp(), sm.Lunch)
	if err != nil {
		log.Println(err)
		return
	}

	meals = todayMeals

}

// /keyboard
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
	log.Println("Server Started")

	feedbackLog, err := os.OpenFile("feedback.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 있으면 사용, 없으면 생성
	if err != nil {
		panic(err)
	}
	feedback = log.New(feedbackLog, "", 0)

	gocron.Every(1).Day().At("00:00").Do(getMeals)

	// 매 시간 16분마다 미세먼지를 불러오게
Sapjil:
	for x := 0; x < 3; x++ {

		for y := 0; y < 10; y++ {
			if x == 2 && y > 3 {
				break Sapjil
			}
			time := fmt.Sprintf("%d%d:16", x, y)
			gocron.Every(1).Day().At(time).Do(getAirq)

		}

	}

	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/keyboard", keyboardHandler)

	// init
	getMeals()
	getAirq()

	go server.ListenAndServe()
	<-gocron.Start()
}
