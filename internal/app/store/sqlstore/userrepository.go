package sqlstore

import "github.com/yashchenkoyurii/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	password, err := u.EncryptedPassword()
	if err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"insert into users (email, password) values ($1, $2) returning id", u.Email, password).
		Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow("select id, email, password from users where email = $1", email).
		Scan(&u.ID, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
