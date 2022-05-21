package main

import (
	"lib-api/handler"
	"log"
	"net/http"
)

func main() {
	handler := handler.NewBookHandler()
	log.Fatal(http.ListenAndServe(":5002", handler))
}
