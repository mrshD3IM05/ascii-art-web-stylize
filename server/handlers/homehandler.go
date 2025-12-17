package handlers

import (
	"net/http"
)

func (a *HandlersStruct) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := a.IndexT
	if r.URL.Path != "/" {
		a.RenderError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		a.RenderError(w, http.StatusBadRequest)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		a.RenderError(w, http.StatusInternalServerError)
	}
}
