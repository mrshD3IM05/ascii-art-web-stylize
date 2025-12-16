package handlers

import (
	"ascii-art-web/server/ascii"
	"bytes"
	"net/http"
)

func (a *HandlersStruct) AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		a.RenderError(w, http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		a.RenderError(w, http.StatusBadRequest)
		return
	}

	asciiResult, status := ascii.Run(text, banner)
	if status != 200 {
		a.RenderError(w, status)
		return
	}

	data := struct {
		Result string
	}{
		Result: asciiResult,
	}

	tmpl := a.ResultT

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		a.RenderError(w, http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
