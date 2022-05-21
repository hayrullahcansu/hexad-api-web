package repo

import (
	"reflect"
	"testing"
)

func TestBookRepository(t *testing.T) {
	repo, err := NewTestBookRepository()
	if err != nil {
		t.Errorf("error when mooking db %q", err)
	}
	t.Run("returns the list of books in the library from db",
		func(t *testing.T) {
			got := repo.GetBooks()
			want := getTestBooks()
			if !reflect.DeepEqual(want, got) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
}
