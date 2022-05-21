package handler

import (
	"lib-api/repo"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGETBooks(t *testing.T) {
	repo, err := repo.NewTestBookRepository()
	if err != nil {
		t.Errorf("cannot initialize test book repository %v", err.Error())
	}
	handler := BookListHandler{repo}
	t.Run("User can view books in library", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/books", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		got := response.Result().StatusCode
		want := 200
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
	t.Run("User can borrow a book from the library", func(t *testing.T) {
		data := url.Values{}
		data.Set("user", "test1_username")
		data.Set("book", "TestBook1")
		request, _ := http.NewRequest(http.MethodPost, "/books/borrow", strings.NewReader(data.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		got := response.Result().StatusCode
		want := 200
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
	t.Run("User can borrow a book from the library", func(t *testing.T) {
		data := url.Values{}
		data.Set("user", "test1_username")
		data.Set("book", "TestBook1")
		request, _ := http.NewRequest(http.MethodPost, "/books/return", strings.NewReader(data.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		got := response.Result().StatusCode
		want := 200
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})

}
