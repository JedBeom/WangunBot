package main

import (
	"log"
	"net/http"
	"time"

	sm "github.com/JedBeom/schoolmeal"
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

	now := time.Now().Local()

	if now.Weekday() == time.Saturday {
		now = now.AddDate(0, 0, 1)
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

	if meals[day].Content == "" {
		meals[day].Content = "없어요!\n왕운봇 소멸의 날이 다가오고 있어요..."
	}

	message := meals[day].Date + "의 급식\n\n" + meals[day].Content
	sendMessage(w, message, weekdays)

}
