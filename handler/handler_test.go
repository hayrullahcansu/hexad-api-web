package handler

import (
	"lib-api/repo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETBooks(t *testing.T) {
	repo, err := repo.NewTestBookRepository()
	if err != nil {
		t.Errorf("cannot initialize test book repository %v", err.Error())
	}
	handler := BookHandler{repo}
	t.Run("returns the list of books in the library", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/books", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		got := response.Result().StatusCode
		want := 200
		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
