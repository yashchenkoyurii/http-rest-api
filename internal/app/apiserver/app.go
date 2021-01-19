package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/yashchenkoyurii/http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

type App struct {
	config *Config
}

func NewApp(config *Config) *App {
	return &App{config: config}
}

func (app *App) Run() error {
	db, err := app.connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	server := newServer(store)

	return http.ListenAndServe(app.config.BindAddr, server)
}

func (app *App) connectDB() (*sql.DB, error) {
	fmt.Println(app.config.Sql.DBUrl())
	db, err := sql.Open("postgres", app.config.Sql.DBUrl())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
