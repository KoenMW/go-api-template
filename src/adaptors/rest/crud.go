package rest

import "net/http"

type CrudHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func RegisterCRUD(path string, handler CrudHandler) {
	http.HandleFunc("POST "+path, handler.Create)
	http.HandleFunc("GET "+path+"/{id}", handler.GetByID)
	http.HandleFunc("GET "+path, handler.List)
	http.HandleFunc("PUT "+path+"/{id}", handler.Update)
	http.HandleFunc("DELETE "+path+"/{id}", handler.Delete)
}
