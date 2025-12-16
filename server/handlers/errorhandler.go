package handlers

import (
	"bytes"
	"log"
	"net/http"
)

func (a *HandlersStruct) RenderError(w http.ResponseWriter, status int) {
	tmpl := a.ErrorT

	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: http.StatusText(status),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Println("Error rendering error template:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)

}
