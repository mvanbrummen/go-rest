package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvanbrummen/go-rest/repository"
)

type TitlesHandler struct {
	router           *mux.Router
	titlesRepository *repository.TitlesRepository
}

func (*TitlesHandler) GetTitle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func NewTitlesHandler(r *mux.Router, titlesRepository *repository.TitlesRepository) {
	handler := &TitlesHandler{
		r,
		titlesRepository,
	}

	r.HandleFunc("/titles/{id}", handler.GetTitle)
}
