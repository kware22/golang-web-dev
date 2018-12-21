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

type menu struct {
	Type        string
	Name        string
	Description string
	Price       float32
}

func main() {

	m1 := menu{
		Type:        "Breakfast",
		Name:        "Eggs N Bacon",
		Description: "Scrambled eggs with your choice of Bacon (pork, turkey)",
		Price:       7.99,
	}
	m2 := menu{
		Type:        "Lunch",
		Name:        "BLT",
		Description: "Bacon, Lettuce and Tomato on a bun",
		Price:       12.99,
	}
	m3 := menu{
		Type:        "Dinner",
		Name:        "Chicken Alfredo",
		Description: "Grilled chicken over alfredo sauce and pasta noodles",
		Price:       15.99,
	}
	menuItems := []menu{m1, m2, m3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", menuItems)
	if err != nil {
		log.Fatalln(err)
	}
}
