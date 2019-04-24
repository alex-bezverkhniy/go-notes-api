package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// App - main app struncture
type App struct {
	Router *mux.Route
	DB     *sql.DB
}

// Initialize - init function
func (a *App) Initialize(user, password, dbname string) {

}

// Run - runs api
func (a *App) Run(address string) {

}
