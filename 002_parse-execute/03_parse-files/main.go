package main

import (
	"html/template"
	"log"
	"os"
)

func checkErr(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	tpl, err := template.ParseFiles("one.gmao")

	checkErr(err)

	err = tpl.Execute(os.Stdout, nil)
	checkErr(err)

	tpl, err = template.ParseFiles("two.gmao", "vespa.gmao")
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	checkErr(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	checkErr(err)

	err = tpl.Execute(os.Stdout, nil)
	checkErr(err)
}
