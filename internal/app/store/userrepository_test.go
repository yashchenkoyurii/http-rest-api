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

	u, err := s.User().Create(model.TestUser(t))

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

	u := model.TestUser(t)
	u.Email = "notfound@email.com"

	_, err := s.User().FinByEmail(u)
	assert.Error(t, err)

	u.Email = "email@email.com"

	s.User().Create(u)
	foundedUser, err := s.User().FinByEmail(u)
	assert.NoError(t, err)
	assert.NotNil(t, foundedUser)

}
