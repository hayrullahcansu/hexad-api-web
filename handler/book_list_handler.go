package handler

import (
	"encoding/json"
	"fmt"
	"lib-api/repo"
	"net/http"
	"regexp"
)

var pattern = `books\/{0,1}(.*)\/{0,1}`

type BookListHandler struct {
	repo.IBookRepository
}

func NewBookListHandler() *BookListHandler {
	db := repo.Instance()
	repo := repo.NewBookRepository(db)
	return &BookListHandler{repo}
}

func (bh *BookListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(pattern)
	url := r.URL.Path[1:]
	if regexGroup := reg.FindStringSubmatch(url); regexGroup != nil {
		bookName := regexGroup[1]
		switch method := r.Method; method {
		case "GET":
			// returns the list of books in the library
			if bookName == "" {
				if data, err := json.Marshal(bh.GetBooks()); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, err.Error())
				} else {
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(data))
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, string(bookName))
			return
		default:
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Invalid URL")
		return
	}
}
