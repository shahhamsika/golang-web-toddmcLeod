package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	mkg := sage{
		"Mohandas Karamchand Gandhi",
		"Non-violence",
	}
	buddha := sage{
		"Gautam Buddha",
		"Meditate",
	}
	christ := sage{
		"Jesus Christ",
		"Son of God",
	}
	sages := []sage{mkg, buddha, christ}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}
}
