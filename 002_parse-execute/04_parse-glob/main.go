package main

import (
	"log"
	"os"
	"text/template"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmp, err := template.ParseGlob("templates/*")
	checkErr(err)

	err = tmp.Execute(os.Stdout, nil)
	checkErr(err)

	err = tmp.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	checkErr(err)

	err = tmp.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	checkErr(err)

	err = tmp.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	checkErr(err)
}
