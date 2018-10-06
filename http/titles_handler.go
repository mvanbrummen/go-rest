package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvanbrummen/go-rest/repository"
)

type TitlesHandler struct {
	router           *mux.Router
	titlesRepository *repository.TitlesRepository
}

func (t *TitlesHandler) GetTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	title, err := t.titlesRepository.FetchTitle(id)
	if err != nil {
		panic(err)
	}

	var b []byte

	if title == nil {
		b = []byte("{}")
	} else {
		b, err = json.Marshal(title)

		if err != nil {
			panic(err)
		}
	}

	w.Write(b)
}

func (t *TitlesHandler) SearchTitle(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.FormValue("q")

	results, err := t.titlesRepository.SearchByTitle(searchTerm, 10)

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(results)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}

func NewTitlesHandler(r *mux.Router, titlesRepository *repository.TitlesRepository) {
	handler := &TitlesHandler{
		r,
		titlesRepository,
	}

	r.HandleFunc("/titles/{id}", handler.GetTitle)
	r.Path("/titles").Queries("q", "{q}").HandlerFunc(handler.SearchTitle)
}
