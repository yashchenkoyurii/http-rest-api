package store

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStore(t *testing.T, config *Config) (*Store, func(...string)) {
	t.Helper()

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}

func Host() (host string) {
	if host = os.Getenv("DB_HOST"); host == "" {
		host = "localhost"
	}

	return
}
