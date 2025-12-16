package handlers

import (
	"bytes"
	"fmt"
	"html/template"
)

func TempParser() (*HandlersStruct, error) {
	IndT, err := template.ParseFiles("templates/index.html")
	ResT, err1 := template.ParseFiles("templates/result.html")
	ErrT, err2 := template.ParseFiles("templates/error.html")
	if err != nil || err1 != nil || err2 != nil {
		return &HandlersStruct{}, fmt.Errorf("template parsing error")
	}
	// Validate templates by executing them with test data
	testRes := struct{ Result string }{Result: "Test"}
	testErr := struct {
		Status  int
		Message string
	}{Status: 200, Message: "OK"}
	
	if !validTmplt(ResT, testRes) || !validTmplt(ErrT, testErr) {
		return &HandlersStruct{}, fmt.Errorf("template validation error")
	}
	
	return &HandlersStruct{
		IndexT:  IndT,
		ResultT: ResT,
		ErrorT:  ErrT,
	}, nil
}

func validTmplt(t *template.Template, testData interface{}) bool {

	var buf bytes.Buffer
	if err := t.Execute(&buf, testData); err != nil {
		return false
	}
	return true
}
