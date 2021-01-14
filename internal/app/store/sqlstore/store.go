package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) store.Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	return &UserRepository{store: s}
}
