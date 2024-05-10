package core

import (
	"log"
	"os"
	"text/template"
)

func Run() {
	renderHome()
	renderAbout()
}

func renderHome() {
	t, err := template.ParseFiles("templates/base.tmpl", "templates/home.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}

func renderAbout() {
	t, err := template.ParseFiles("templates/base.tmpl", "templates/about.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}
