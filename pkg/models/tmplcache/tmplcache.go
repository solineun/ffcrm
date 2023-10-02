package tmplcache

import (
	"html/template"
)

type TemplateCache struct {
	strg map[string]*template.Template
}

func (c *TemplateCache) GetTemplate(name string) (*template.Template, bool) {
	t, ok := c.strg[name]
	return t, ok
}

func (c *TemplateCache) InsertTemplate(name string, t *template.Template) {
	c.strg[name] = t
}

func NewTemplateCache(dir string) *TemplateCache {
	return &TemplateCache{
		strg: make(map[string]*template.Template),
	}
}