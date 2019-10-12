package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tp.gohtml")

	if err != nil {
		log.Fatal("Failed to create file", err)
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatal(err)
	}

}
