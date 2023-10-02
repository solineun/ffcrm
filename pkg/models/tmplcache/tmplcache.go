package tmplcache

import (
	"html/template"
	"path/filepath"
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

func NewTemplateCache(dir string) (*TemplateCache, error){
	cache := &TemplateCache{
		strg: make(map[string]*template.Template),
	}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
    if err != nil {
        return nil, err
    }
 
    for _, page := range pages {
        name := filepath.Base(page)
 
        ts, err := template.ParseFiles(page)
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
 
        cache.InsertTemplate(name, ts)
	}

	return cache, nil

}