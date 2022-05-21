package handler

import (
	"encoding/json"
	"fmt"
	"lib-api/repo"
	"net/http"
	"regexp"
)

var patternBorrowed = `borrowed\/{0,1}(.*)`

type BorrowedListHandler struct {
	repo.ILibraryRepository
}

func NewBorrowedListHandler() *BorrowedListHandler {
	db := repo.Instance()
	repo := repo.NewLibraryRepository(db)
	return &BorrowedListHandler{repo}
}

func (bh *BorrowedListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(patternBorrowed)
	url := r.URL.Path[1:]
	if regexGroup := reg.FindStringSubmatch(url); regexGroup != nil {
		user := regexGroup[1]
		switch method := r.Method; method {
		case "GET":
			// returns borrowed list
			if user != "" {
				if data, err := json.Marshal(bh.MyBorrowedList(user)); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, err.Error())
				} else {
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(data))
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, string(user))
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
