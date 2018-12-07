package main

import (
	"log"
	"net/http"
	"time"

	sm "github.com/jedbeom/schoolmeal"
)

var (
	school = sm.School{
		SchoolCode:     "Q100005451",
		SchoolKindCode: sm.Middle,
		Zone:           sm.Jeonnam,
	}

	// 급식 저장용
	meals []sm.Meal

	// feedback.log
	feedback *log.Logger
)

// 급식을 불러옴
func getMeals() {

	now := time.Now()

	if now.Weekday() == time.Saturday {
		now.Add(time.Hour * 24)
	}
	todayMeals, err := school.GetWeekMeal(sm.Timestamp(now), sm.Lunch)
	if err != nil {
		log.Println(err)
		return
	}

	meals = todayMeals

}

func sendMeal(w http.ResponseWriter, post Post) {

	var day int
	switch post.Content {
	case "월요일":
		day = 1
	case "화요일":
		day = 2
	case "수요일":
		day = 3
	case "목요일":
		day = 4
	case "금요일":
		day = 5
	}

	message := meals[day].Date + "\n" + meals[day].Content
	sendMessage(w, message, weekdays)

}
