package memorystore

import (
	"github.com/yashchenkoyurii/http-rest-api/internal/app/model"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store"
)

type memoryStore map[string]*model.User

type Store struct {
	users memoryStore
}

func New() *Store {
	return &Store{users: make(memoryStore)}
}

func (s *Store) User() store.UserRepository {
	return &UserRepository{users: s.users}
}
