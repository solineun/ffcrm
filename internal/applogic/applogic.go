package applogic

import (
	"html/template"
	"net/http"
	"path/filepath"

	tmpl "github.com/solineun/ffcrm/internal/templates"
)

type Application struct {
	tmplcache tmpl.TemplateCacheStorage
	
}

func (app *Application) UpdateTmplCache(dir string) error {
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
    if err != nil {
        return err
    }
 
    for _, page := range pages {
        name := filepath.Base(page)
 
        ts, err := template.ParseFiles(page)
        if err != nil {
            return err
        }

        ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
        if err != nil {
            return err
        }
 
        ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
        if err != nil {
            return err
        }
 
        app.tmplcache.InsertTemplate(name, ts)
	}

	return nil
}

func (app *Application) RenderTmpl(w http.ResponseWriter, r *http.Request, name string, td *tmpl.TemplateData) {
	ts, ok := app.tmplcache.GetTemplate(name)
	if !ok {
		//LOG SERVER ERR
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		//LOG SERVER ERR 
	}
}