package handlers

import (
	"net/http"
)

func (a *HandlersStruct) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.RenderError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		a.RenderError(w, http.StatusBadRequest)
		return
	}

	if err := a.IndexT.Execute(w, nil); err != nil {
		a.RenderError(w, http.StatusInternalServerError)
	}
}
