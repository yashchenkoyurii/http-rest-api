package sqlstore

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDB(t *testing.T) (*sql.DB, func(...string)) {
	c := NewConfig()
	c.Host = os.Getenv("DB_HOST")
	db, err := sql.Open("postgres", c.DBUrl())

	if err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", ")))
		}
	}
}
