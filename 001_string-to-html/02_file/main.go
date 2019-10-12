package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Hamsika Shah"

	tp := `
	<html>
	<head>
	<meta charset="UTF-8">
	<title> Hello world </title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body> </html>`

	fmt.Println(tp)

	file, err := os.Create("index.html")

	if err != nil {
		log.Fatal("Failed to create File", err)
	}

	defer file.Close()

	io.Copy(file, strings.NewReader(tp))

}
