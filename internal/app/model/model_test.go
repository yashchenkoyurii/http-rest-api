package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/model"
	"testing"
)

func TestUser_EncryptedPassword(t *testing.T) {
	u := model.TestUser(t)
	password, _ := u.EncryptedPassword()
	assert.NotEmpty(t, password)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "Valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "Empty email",
			u: func() (u *model.User) {
				u = model.TestUser(t)
				u.Email = ""
				return
			},
			isValid: false,
		},
		{
			name: "Empty password",
			u: func() (u *model.User) {
				u = model.TestUser(t)
				u.Password = ""
				return
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}

}
