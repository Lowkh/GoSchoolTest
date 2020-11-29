package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"tawesoft.co.uk/go/dialog"
)

var tpl *template.Template

func init() {
	//Path of templates
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":5050", nil) //Setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		gotError := false
		operation := req.FormValue("action")
		a, err := strconv.Atoi(req.FormValue("first_number"))
		if err != nil {
			dialog.Alert("Error: %v", err)
			gotError = true
		}
		b, err := strconv.Atoi(req.FormValue("sec_number"))
		if err != nil {
			dialog.Alert("Error: %v", err)
			gotError = true
		}
		switch {
		case gotError:

		case operation == "add":
			dialog.Alert("Answer: %v", Add(a, b))
		case operation == "subtract":
			dialog.Alert("Answer: %v", Subtract(a, b))
		case operation == "multiply":
			dialog.Alert("Answer: %v", Multiply(a, b))
		case operation == "divide":
			dialog.Alert("Answer: %v", Divide(a, b))
		}
	}

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
