package templates

import (
	"html/template"

	"github.com/solineun/ffcrm/pkg/models"
)

type TemplateData struct {
	Order    *models.Order
	LastFive []*models.Order
}

type TemplateCacheStorage interface {
	GetTemplate(name string) (*template.Template, bool)
	InsertTemplate(name string, t *template.Template)
}
