package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETBooks(t *testing.T) {
	t.Run("returns the list of books in the library", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/books", nil)
		response := httptest.NewRecorder()
		BookList(response, request)
		got := response.Result().StatusCode
		want := 200
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
