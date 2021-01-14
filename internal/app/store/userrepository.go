package store

import "github.com/yashchenkoyurii/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	password, err := u.EncryptedPassword()
	if err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"insert into users (email, password) values ($1, $2) returning id", u.Email, password).
		Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FinByEmail(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow("select id, email, password from users where email = $1", u.Email).
		Scan(&u.ID, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
