package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
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
	bookListHandler := NewBookListHandler()
	a.mux.HandleFunc("/api/v1/books/", bookListHandler.ServeHTTP)
	borrowedListHandler := NewBorrowedListHandler()
	a.mux.HandleFunc("/api/v1/borrowed/", borrowedListHandler.ServeHTTP)

	corsHandler := cors.Default().Handler(a.mux)
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	uri := ":" + port
	http.ListenAndServe(uri, corsHandler)
}
