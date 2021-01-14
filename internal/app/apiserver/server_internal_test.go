package apiserver

import (
	"github.com/stretchr/testify/assert"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store/memorystore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s := newServer(memorystore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
