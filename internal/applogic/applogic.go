package applogic

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/solineun/ffcrm/internal/logger"
	"github.com/solineun/ffcrm/internal/sql"
	tmpl "github.com/solineun/ffcrm/internal/templates"
)

type Application struct {
	tmplcache tmpl.TemplateCacheStorage
	errLog logger.ErrLogger
	db sql.FfcrmDb
}

func NewApplication(tmplcache tmpl.TemplateCacheStorage, 
					errLog logger.ErrLogger, 
					db sql.FfcrmDb) *Application {						
	return &Application{
		tmplcache: tmplcache,
		errLog: errLog,
		db: db,
	}
}

func (app *Application) RenderTmpl(w http.ResponseWriter, r *http.Request, name string, td *tmpl.TemplateData) {
	ts, ok := app.tmplcache.GetTemplate(name)
	if !ok {
		app.errLog.ServerError(w, fmt.Errorf("no %s template was found", name))
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		app.errLog.ServerError(w, err) 
	}
}


type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/order", app.showOrder)
	mux.HandleFunc("/order/create", app.createOrder)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}