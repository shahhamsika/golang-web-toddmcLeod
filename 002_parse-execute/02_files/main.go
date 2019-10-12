package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tp.gohtml")

	if err != nil {
		log.Fatal("Failed to create file", err)
	}

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = tpl.Execute(file, nil)

	if err != nil {
		log.Fatal(err)
	}

}
