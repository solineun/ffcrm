package main

import (
	"html/template"

	"github.com/solineun/ffcrm/pkg/models"
)

type templateData struct {
	Order *models.Order
	LastFive []*models.Order
}

type CacheStorage interface {
	GetTemplate(name string) (*template.Template, bool)
	InsertTemplate(name string, t *template.Template)
}