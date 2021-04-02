package auth

import (
	"database/sql"

	"github.com/gorilla/mux"
	"sirlana.com/sirlana/sso/libs"
)

type Auth struct {
	r *mux.Router
}

var (
	jwt *libs.JWT
	db  *sql.DB
)

func NewAuth(j *libs.JWT, r *mux.Router, d *sql.DB) *Auth {
	jwt = j
	db = d
	return &Auth{
		r: r,
	}
}

func (a Auth) Run() {
	auth := a.r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signin", signIn).Methods("POST")
	auth.HandleFunc("/signup", signUp).Methods("POST")
}
