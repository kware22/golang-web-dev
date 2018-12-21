package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl2.gohtml"))
}
func main() {

	sages := map[string]string{
		"Omni-Present": "Jesus",
		"USA":          "MLK",
		"US":           "Malcolm X",
		"U.S.A":        "Shirley Chisolm",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
