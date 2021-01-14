package sqlstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/model"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := TestDB(t)
	defer teardown("users")

	s := New(db)
	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := TestDB(t)
	defer teardown("users")

	s := New(db)
	u := model.TestUser(t)
	u.Email = "notfound@email.com"

	_, err := s.User().FindByEmail(u.Email)
	assert.Error(t, err)

	u.Email = "email@email.com"

	s.User().Create(u)
	foundedUser, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, foundedUser)
}
