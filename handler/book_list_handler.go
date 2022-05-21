package handler

import (
	"encoding/json"
	"fmt"
	"lib-api/repo"
	"net/http"
	"regexp"
)

var pattern = `books\/{0,1}(borrow|return)?`

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
		action := regexGroup[1]
		switch method := r.Method; method {
		case "GET":
			// returns the list of books in the library
			if action == "" {
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
			fmt.Fprint(w, string(action))
			return
		case "POST", "PUT":
			err := r.ParseForm()
			if err == nil && len(r.Form) > 1 && action != "" {
				user := r.Form.Get("user")
				book := r.Form.Get("book")
				borrow, err := bh.BorrowOrReturnBook(action, user, book)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, err.Error())
					return
				}
				if data, err := json.Marshal(borrow); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, err.Error())
				} else {
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(data))
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Invalid Request")
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Unsupported Method")
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Invalid URL")
		return
	}
}
