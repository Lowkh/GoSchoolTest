package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	//Path of templates
	//tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := http.ListenAndServe(":5050", handlers()) //Setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handlers() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/", index)
	h.HandleFunc("/test", test)
	return h
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		gotError := false
		operation := req.FormValue("action")
		a, err := strconv.Atoi(req.FormValue("first_number"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(res, "missing value")
			gotError = true
		}
		b, err := strconv.Atoi(req.FormValue("sec_number"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(res, "missing value")
			gotError = true
		}
		switch {
		case gotError:

		case operation == "add":
			//dialog.Alert("Answer: %v", Add(a, b))
			fmt.Fprintln(res, Add(a, b))
		case operation == "subtract":
			//dialog.Alert("Answer: %v", Subtract(a, b))
			fmt.Fprintln(res, Subtract(a, b))
		case operation == "multiply":
			//dialog.Alert("Answer: %v", Multiply(a, b))
			fmt.Fprintln(res, Multiply(a, b))
		case operation == "divide":
			//dialog.Alert("Answer: %v", Divide(a, b))
			fmt.Fprintln(res, Divide(a, b))
		}
	}

	//tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func test(res http.ResponseWriter, req *http.Request) {
	//tpl.ExecuteTemplate(res, "test.gohtml", nil)
}
