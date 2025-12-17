package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type PageData struct {
	Status  int
	Message string
	Result  string
}

func TempParser() (*HandlersStruct, error) {
	errT := 0
	T, err := template.ParseFiles("templates/index.html", "templates/error.html", "templates/result.html")
	if err != nil {
		log.Print(err)
		errT++
	}
	// Validate templates by executing them with test data
	tests := []struct {
		page string
		data *PageData
	}{
		{
			page: "index",
			data: nil,
		},
		{
			page: "result",
			data: &PageData{Result: "Test"},
		},
		{
			page: "error",
			data: &PageData{Status: 200, Message: "OK"},
		},
	}
	valids := make(map[string]*template.Template)
	for _, t := range tests {
		tmpl := T.Lookup(t.page + ".html")
		if tmpl == nil {
			log.Printf("template %s not found", t.page)
			errT++
		}
		if !validTmplt(tmpl, t.data) {
			log.Printf("invalid templates placeholder in \"%s\"", t.page)
			errT++
		}
		valids[t.page] = tmpl
	}
	if errT != 0 {
		return nil, fmt.Errorf("check parsing errors")
	}
	return &HandlersStruct{
		IndexT:  valids["index"],
		ResultT: valids["result"],
		ErrorT:  valids["error"],
	}, nil
}

func validTmplt(t *template.Template, testData any) bool {
	var buf bytes.Buffer
	if err := t.Execute(&buf, testData); err != nil {
		return false
	}
	return true
}
