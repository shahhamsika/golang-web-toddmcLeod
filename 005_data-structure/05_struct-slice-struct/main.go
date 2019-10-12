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

type car struct {
	Model  string
	Milage float64
}

type item struct {
	Sages []sage
	Cars  []car
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

	c1 := car{
		"SUV",
		41.25,
	}
	c2 := car{
		"Nano",
		4.25,
	}
	cars := []car{c1, c2}

	data := item{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
