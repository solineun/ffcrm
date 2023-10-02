package applogic

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	tmpl "github.com/solineun/ffcrm/internal/templates"
	"github.com/solineun/ffcrm/pkg/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errLog.ClientError(w, http.StatusNotFound)
		return
	}

	orders, err := app.db.LatestFiveOrders()
	if err !=  nil {
		if !errors.Is(err, models.ErrNoRecord) {
			app.errLog.ServerError(w, err)
			return
		}
	}
	
	app.RenderTmpl(w, r, "home.page.tmpl", &tmpl.TemplateData{
		LastFive: orders,
	})
}

func (app *Application) showOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.errLog.ClientError(w, http.StatusNotFound)
		return
	}
	
	order, err := app.db.GetOrderById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.errLog.ClientError(w, http.StatusNotFound)
		} else {
			app.errLog.ServerError(w, err)
		}
		return
	}

	app.RenderTmpl(w, r, "show.page.tmpl", &tmpl.TemplateData{
		Order: order,
	})
}

func (app *Application) createOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.errLog.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	productName := "test"
	id, err := app.db.InsertOrder(productName)
	if err != nil {
		if errors.Is(err, models.ErrLongValue) {
			app.errLog.ClientError(w, http.StatusBadRequest)
			return
		}
		app.errLog.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/order?id=%d", id), http.StatusSeeOther)
}