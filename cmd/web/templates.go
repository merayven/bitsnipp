package main

import (
	"bitsnipp.merayven.net/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
