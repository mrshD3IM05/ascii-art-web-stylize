package handlers

import (
	"html/template"
	"net/http"
)

type HandlersStruct struct {
	IndexT  *template.Template
	ResultT *template.Template
	ErrorT  *template.Template
}

func Handlers() {

	H, errT := TempParser()
	if errT != nil {
		panic("Failed to parse templates")
	}

	http.HandleFunc("/", H.HomeHandler)
	http.HandleFunc("/ascii-art", H.AsciiHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("templates")),
		),
	)
}
