package handler

import "net/http"

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type orderHandler interface {
	ShowOrder(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type HandlerAdapter struct {
	order orderAdapter
}

func (h HandlerAdapter) Home(w http.ResponseWriter, r *http.Request) {
	
}

type orderAdapter struct {

}