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
	repo, err := repo.NewTestLibraryRepository()
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
		data.Set("User", "test1_username")
		data.Set("Book", "TestBook1")
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
	borrowedHandler := BorrowedListHandler{repo}

	t.Run("User can see borrowed list", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/borrowed/test1_username", nil)
		response := httptest.NewRecorder()
		borrowedHandler.ServeHTTP(response, request)
		got := response.Result().StatusCode
		want := 200
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("User can return a book from the library", func(t *testing.T) {
		data := url.Values{}
		data.Set("User", "test1_username")
		data.Set("Book", "TestBook1")
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
