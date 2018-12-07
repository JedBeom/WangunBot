package main

import (
	"bytes"
	"net/http"
	"text/template"
	"time"
)

type Event struct {
	Name string
	Date time.Time
	DDay int
}

var (
	Events []Event
)

func init() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	Events = []Event{
		Event{
			Date: time.Date(2018, 12, 18, 0, 0, 0, 0, loc),
			Name: "학생회장선거",
		},
		Event{
			Date: time.Date(2018, 12, 26, 0, 0, 0, 0, loc),
			Name: "왕운축제",
		},
		Event{
			Date: time.Date(2018, 12, 31, 0, 0, 0, 0, loc),
			Name: "방학식",
		},
	}

}

func dDay(w http.ResponseWriter) {
	for i := range Events {
		Events[i].DDay = int(Events[i].Date.Sub(time.Now()).Hours() / 24)
	}

	format := `주요 일정
++++++++++

{{ range . }}
{{ .Name }} {{ .Date }}
D-{{ .DDay }}
---
{{ end }}`
	t := template.Must(template.New("format").Parse(format))

	var tpl bytes.Buffer

	t.Execute(&tpl, Events)
	sendMessage(w, tpl.String(), home)

}
