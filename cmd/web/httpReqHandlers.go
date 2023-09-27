package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"solineun/ffcrm/pkg/models"
	"strconv"
)

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

	for _, o := range orders {
		fmt.Fprint(w, o.Format())
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
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

	fmt.Fprint(w, order.Format())
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	productName := "test"
	id, err := app.orders.Insert(productName)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/order?id=%d", id), http.StatusSeeOther)
}