package rest

import (
	"go-api/ports/handler"
	"net/http"
)

func RegisterCRUD(mux *http.ServeMux, path string, handler handler.Basehandler) {
	mux.HandleFunc("POST "+path, handler.Create)
	mux.HandleFunc("GET "+path+"/{id}", handler.GetByID)
	mux.HandleFunc("GET "+path, handler.List)
	mux.HandleFunc("PUT "+path+"/{id}", handler.Update)
	mux.HandleFunc("DELETE "+path+"/{id}", handler.Delete)
}
