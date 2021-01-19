package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func newServer(store store.Store) *server {
	s := &server{router: mux.NewRouter(), logger: logrus.New(), store: store}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.HandleUsersCreate()).Methods("POST")
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//
	}
}
