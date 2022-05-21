package main

import "lib-api/handler"

func main() {
	apiListener := handler.NewApiListener()
	apiListener.ListenAndServe()
}
