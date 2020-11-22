package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		s.config.Host, s.config.DBPort, s.config.User, s.config.Password, s.config.DBName, s.config.SSLMode)
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	return &UserRepository{store: s}
}
