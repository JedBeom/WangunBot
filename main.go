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
	gocron.Every(1).Hour().Do(getAirq)

	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/keyboard", keyboardHandler)

	getMeals()
	getAirq()

	go server.ListenAndServe()
	<-gocron.Start()
}
