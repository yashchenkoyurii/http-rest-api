package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func (u *User) EncryptedPassword() (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 30)),
	)

	return err
}
