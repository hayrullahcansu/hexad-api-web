package repo

import (
	"testing"
)

func TestDbContext(t *testing.T) {
	_, err := getDbContext()
	if err != nil {
		t.Errorf("error db context %q", err)
	}
}
