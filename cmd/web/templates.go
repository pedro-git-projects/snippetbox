package main

import (
	"html/template"
	"path/filepath"

	"github.com/pedro-git-projects/snippetbox/internal/models"
)

// will store the dynamic data needed for template interpolation
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = templateSet
	}
	return cache, nil
}
