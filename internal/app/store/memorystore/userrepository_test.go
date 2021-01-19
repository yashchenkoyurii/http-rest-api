package memorystore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/model"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store/memorystore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := memorystore.New()
	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := memorystore.New()
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
