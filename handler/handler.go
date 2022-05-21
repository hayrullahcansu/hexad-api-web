package handler

import (
	"encoding/json"
	"fmt"
	"lib-api/repo"
	"net/http"
)

type BookHandler struct {
	repo.IBookRepository
}

func NewBookHandler() *BookHandler {
	db := repo.Instance()
	repo := repo.NewBookRepository(db)
	return &BookHandler{repo}
}

func (bh *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if data, err := json.Marshal(bh.GetBooks()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(data))
	}
}
