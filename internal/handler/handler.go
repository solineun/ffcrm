package handler

import "net/http"

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
	orderHandler
}

type orderHandler interface {
	ShowOrder(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
}