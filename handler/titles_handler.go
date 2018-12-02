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
	titlesRepository repository.ITitlesRepository
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

	if limitParam == "" {
		limitParam = "10"
	}

	var err error
	limit, err := strconv.Atoi(limitParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"limit must be an integer"}`)
		return
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

func NewTitlesHandler(titlesRepository repository.ITitlesRepository) *TitlesHandler {
	return &TitlesHandler{
		titlesRepository,
	}
}
