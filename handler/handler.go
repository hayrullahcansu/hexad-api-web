package handler

import (
	"encoding/json"
	"fmt"
	"lib-api/data"
	"net/http"
)

func BookList(w http.ResponseWriter, r *http.Request) {
	if data, err := json.Marshal(GetBooks()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, data)
	}
}
func GetBooks() []data.Book {
	return make([]data.Book, 0)
}
