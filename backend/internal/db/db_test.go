package db_test

import (
	"testing"

	"github.com/Retasusan/colab_backend/internal/db"
)

func TestOpen(t *testing.T) {
	conn, err := db.Open()
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		t.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
