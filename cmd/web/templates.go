package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/rufusbarnes/prettyBookmarks/pkg/forms"
)

type templateData struct {
	CSRFToken string
	Flash     string
	Form      *forms.Form
}

func humanDate(t time.Time) string {
	return t.Format("2 Jan 2006 at 15:04pm")
}

// Each function must return either (1 value) or (1 value, err)
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
