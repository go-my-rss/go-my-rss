package database

import "testing"

func TestDB(t *testing.T) {
	if DB == nil {
		t.Fatal("DB is nil, init failed")
	}
}
