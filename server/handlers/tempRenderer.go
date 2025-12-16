package handlers

import (
	"fmt"
	"html/template"
)

func TempParser() (*HandlersStruct, error) {
	IndT, err := template.ParseFiles("templates/index.html")
	ResT, err1 := template.ParseFiles("templates/result.html")
	ErrT, err2 := template.ParseFiles("templates/error.html")
	if err != nil || err1 != nil || err2 != nil {
		return &HandlersStruct{
			IndexT:  nil,
			ResultT: nil,
			ErrorT:  nil,
		}, fmt.Errorf("template parsing error")
	}
	return &HandlersStruct{
		IndexT:  IndT,
		ResultT: ResT,
		ErrorT:  ErrT,
	}, nil
}
