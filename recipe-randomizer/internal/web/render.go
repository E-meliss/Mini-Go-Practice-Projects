package web

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/recipe.html
var recipePage string

var tpl = template.Must(template.New("recipe").Parse(recipePage))

func RenderRecipePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Safety guard: if embed failed or file missing, show a helpful message.
	if recipePage == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`<pre>Template asset not embedded.
Expected file at: internal/web/templates/recipe.html
Check the path and filename, then re-run: go run ./cmd/api</pre>`))
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, "template render error: "+err.Error(), http.StatusInternalServerError)
	}
}
