package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	limitParam := r.FormValue("limit")

	var limit int
	if limitParam == "" {
		limit = 10
	} else {
		var err error
		limit, err = strconv.Atoi(limitParam)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error":"limit must be an integer"}`)
			return
		}
	}

	results, err := t.titlesRepository.SearchByTitle(searchTerm, limit)

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
	r.Path("/titles").Queries("q", "{q}", "limit", "{limit}").HandlerFunc(handler.SearchTitle)
}
