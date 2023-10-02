package controller

import "net/http"

type ProductsCtrlRepository interface {
	ShowAll(w http.ResponseWriter, r *http.Request)
	ShowById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}
