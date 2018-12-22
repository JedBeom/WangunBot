package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Event struct {
	Name string

	Date          time.Time
	DateTimestamp string

	DDay int
}

var (
	Events []Event
	Loc    *time.Location
)

func init() {
	Loc, _ = time.LoadLocation("Asia/Seoul")

	//newEvent("학생회장선거", 2018, 12, 18)
	newEvent("왕운축제", 2018, 12, 26)
	newEvent("방학식", 2018, 12, 31)

}

func newEvent(name string, year, month, day int) {
	event := Event{
		Name: name,
		Date: time.Date(year, time.Month(month), day, 0, 0, 0, 0, Loc),
	}
	event.DateTimestamp = fmt.Sprintf("%d-%d", month, day)
	Events = append(Events, event)
}

func dDay(w http.ResponseWriter) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, Loc)

	for i := range Events {
		Events[i].DDay = int(Events[i].Date.Sub(today).Hours() / 24)
	}

	format := `{{ range . }}
{{ .DateTimestamp }} {{ .Name }} 
D-{{ .DDay }}
{{ end }}`

	t := template.Must(template.New("format").Parse(format))

	var tpl bytes.Buffer

	t.Execute(&tpl, Events)
	sendMessage(w, tpl.String(), home)

}
