package app

import (
	"database/sql"
	"net/http"

	"sirlana.com/sirlana/sso/libs"

	"github.com/gorilla/mux"
)

var (
	jwt    *libs.JWT
	jwtKey = "for(sirlana 309){}"
)

func Run(r *mux.Router, db *sql.DB) {
	// Init JWT
	jwt = libs.NewJWT(jwtKey)
	jwt.AddExpiredDate(720)

	// Define middleware
	r.Use(Middleware)

	// Init services
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("X-Session-Token")
		if tkn, claims, err := jwt.Decode(token); err != nil && tkn.Valid {
			if !jwt.IsExpired(claims["exp"].(float64)) {
				next.ServeHTTP(w, r)
			} else {
				println("Token Invalid")
			}
		} else {
			println("Token invalid.")
		}
	})
}
