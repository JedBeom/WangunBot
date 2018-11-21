package main

type Post struct {
	UserKey string `json:"user_key"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Message struct {
	Text string `json:"text"`
}

type Keyboard struct {
	Type    string   `json:"type"`
	Buttons []string `json:"buttons"`
}

type Response struct {
	Message  `json:"message"`
	Keyboard `json:"keyboard"`
}

var (
	// '급식'을 받았을 때 keyboard
	weekdays = []string{
		"월요일",
		"화요일",
		"수요일",
		"목요일",
		"금요일",
		"홈",
	}

	// 평상시 keyboard
	buttons = []string{
		"급식",
		"미세먼지",
		"피드백",
	}
)
