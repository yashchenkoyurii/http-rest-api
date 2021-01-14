package memorystore

import (
	"errors"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/model"
)

type UserRepository struct {
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]

	if !ok {
		return nil, errors.New("not found")
	}

	return u, nil
}
