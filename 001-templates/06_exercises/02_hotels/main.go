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

type caliHotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func main() {

	h1 := caliHotel{
		Name:    "Marriott",
		Address: "1234 Main Street",
		City:    "San Francisco",
		Zip:     123456,
		Region:  "Central",
	}
	h2 := caliHotel{
		Name:    "Sheraton",
		Address: "6785 West Ave",
		City:    "Los Angelos",
		Zip:     123890,
		Region:  "Southern",
	}

	hotels := []caliHotel{h1, h2}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
