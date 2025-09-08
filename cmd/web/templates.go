// Holds dynamic data that would be passed to HTML templates.
package main

import "github.com/VicTheM/snippetbox/internal/models"

type templateData struct {
	Snippet *models.Snippet
}
