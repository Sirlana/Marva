package auth

import (
	"encoding/json"
	"net/http"
)

type SignIn struct {
	UserID       int
	Email        string
	Username     string
	FirstName    string
	LastName     string
	ProfilePhoto string
}

func signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usernameOrEmail := r.FormValue("Username-Or-Email")
	password := r.FormValue("Password")
	productID := r.FormValue("Product-Id")
	row := db.QueryRow(`SELECT user_id, status_id, email, username, first_name, last_name, profile_photo 
		FROM users 
		WHERE username=$1 OR email=$2 AND password=$3`,
		usernameOrEmail, usernameOrEmail, password)
	var userID, statusID int
	var email, username, firstName, lastName, profilePhoto string
	if err := row.Scan(&userID, &statusID, &email, &username, &firstName, &lastName, &profilePhoto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// status id 1 is ACTIVE
	if statusID == 1 {
		if _, err := db.Exec("INSERT INTO activities (user_id, product_id, activity) VALUES($1, $2, $3)", userID, productID, "SIGNIN"); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}
	signin := SignIn{
		UserID:       userID,
		Email:        email,
		Username:     username,
		FirstName:    firstName,
		LastName:     lastName,
		ProfilePhoto: profilePhoto,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(signin)
}
