package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]

	tp := `
	<html>
	<head>
	<meta charset="UTF-8">
	<title> Hello world </title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body> </html>`

	file, err := os.Create("index.html")

	if err != nil {
		log.Fatal("File creation failed", err)
	}

	defer file.Close()

	io.Copy(file, strings.NewReader(tp))
}
