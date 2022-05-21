package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ApiListener struct {
	mux *http.ServeMux
}

// NewApiListener return new instance of ApiWorker
func NewApiListener() *ApiListener {
	return &ApiListener{
		mux: http.NewServeMux(),
	}
}

func (a *ApiListener) ListenAndServe() {

	a.mux.HandleFunc("/", TestServer)
	bookListHandler := NewBookListHandler()
	a.mux.HandleFunc("/api/v1/books/", bookListHandler.ServeHTTP)
	borrowedListHandler := NewBorrowedListHandler()
	a.mux.HandleFunc("/api/v1/borrowed/", borrowedListHandler.ServeHTTP)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	uri := ":" + port
	http.ListenAndServe(uri, a.mux)
}

func TestServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello, %s!", r.URL.Path[1:])

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
