package main

import "github.com/pedro-git-projects/snippetbox/internal/models"

// will store the dynamic data needed for template interpolation
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
