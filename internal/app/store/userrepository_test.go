package store

import (
	"github.com/stretchr/testify/assert"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/model"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, &Config{
		Host:     Host(),
		DBName:   "restapi",
		DBPort:   5432,
		User:     "postgres",
		Password: "mysecretpassword",
		SSLMode:  "disable",
	})
	defer teardown("users")

	u, err := s.User().Create(model.User{
		Email:    "example@email.com",
		Password: "test",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FinByEmail(t *testing.T) {
	s, teardown := TestStore(t, &Config{
		Host:     Host(),
		DBName:   "restapi",
		DBPort:   5432,
		User:     "postgres",
		Password: "mysecretpassword",
		SSLMode:  "disable",
	})
	defer teardown("users")

	_, err := s.User().FinByEmail(model.User{Email: "notfound@email.com"})
	assert.Error(t, err)

	email := "email@email.com"

	s.User().Create(model.User{Email: email})
	u, err := s.User().FinByEmail(model.User{Email: email})
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
