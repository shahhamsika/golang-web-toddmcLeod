package main

import "fmt"

func main(){
	name := "Hamsika Shah"

	tp := `
	<html>
	<head>
	<meta charset="UTF-8">
	<title> Hello world </title>
	</head>
	<body>
	<h1>`+ name + `</h1>
	</body> </html>`

	fmt.Println(tp)
}

//run as :- go run main.go > h1.html