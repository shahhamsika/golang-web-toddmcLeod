package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	buddha := sage{
		"Gautam Buddha",
		"Meditate",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", buddha)
	if err != nil {
		log.Fatal(err)
	}
}

// tpl2.gohtml
