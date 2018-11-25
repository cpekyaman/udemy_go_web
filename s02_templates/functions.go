package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type Sage struct {
	Name  string
	Motto string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"uc":  strings.ToUpper,
	"ft":  firstThree,
	"dmy": dayMonthYear,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("functions.gohtml").Funcs(fm).ParseFiles("tpl/functions.gohtml"))
}

func main() {
	buddha := Sage{
		Name:  "Buddha",
		Motto: "blah blah",
	}

	gandhi := Sage{
		Name:  "Gandhi",
		Motto: "ho ho he",
	}

	data := struct {
		Title string
		Sages []Sage
		Time  time.Time
	}{
		Title: "List of Sages",
		Sages: []Sage{buddha, gandhi},
		Time:  time.Now(),
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
