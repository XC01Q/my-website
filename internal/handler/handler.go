package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Handler struct {
	templates *template.Template
	templDir  string
}

func New(templDir string) *Handler {
	return &Handler{templDir: templDir}
}

func (h *Handler) LoadTemplates() error {
	tmpl, err := template.ParseGlob(filepath.Join(h.templDir, "*.html"))
	if err != nil {
		return err
	}
	h.templates = tmpl
	return nil
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.templates.ExecuteTemplate(w, "index.html", nil)
}
