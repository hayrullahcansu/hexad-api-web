package main

import (
	"lib-api/handler"
	"lib-api/repo"
)

func main() {
	repo.SeedDb()
	apiListener := handler.NewApiListener()
	apiListener.ListenAndServe()
}
