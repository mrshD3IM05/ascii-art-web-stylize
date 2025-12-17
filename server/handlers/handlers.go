package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type HandlersStruct struct {
	IndexT  *template.Template
	ResultT *template.Template
	ErrorT  *template.Template
}

func Handlers() {

	H, errT := TempParser()
	if errT != nil {
		log.Printf("Failed to parse templates: %s", errT)
		os.Exit(0)
	}

	http.HandleFunc("/", H.HomeHandler)
	http.HandleFunc("/ascii-art", H.AsciiHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("templates")),
		),
	)
}
