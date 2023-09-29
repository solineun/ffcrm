package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"solineun/ffcrm/pkg/models"
	"strconv"
)

type templateData struct {
	Order *models.Order
	LastFive []*models.Order
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	orders, err := app.orders.Latest()
	if err !=  nil {
		if !errors.Is(err, models.ErrNoRecord) {
			app.serverError(w, err)
			return
		}
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	
	if err = renderTemplate(w, files, orders); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	
	order, err := app.orders.Get(&id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := templateData{Order: order}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	if err = renderTemplate(w, files, data); err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	productName := "**********************************************************************"
	id, err := app.orders.Insert(productName)
	if err != nil {
		if errors.Is(err, models.ErrLongValue) {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/order?id=%d", id), http.StatusSeeOther)
}

func renderTemplate(w io.Writer, files []string, data any) error {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	err = ts.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}