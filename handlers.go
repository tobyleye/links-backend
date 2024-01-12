/*
	endpoints

	POST /login
	POST /register
	POST /createlinks
	GET /listlinks
	GET /profile
	PATCH /profile/update

*/

package main

import (
	"encoding/json"
	"net/http"
)

type Link struct {
	Platform string
	URL      string
}

// ---------- handlers -----

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	type LoginBody struct {
		Email    string
		Password string
	}
	var credentials LoginBody
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	user, err := FindUserByEmail(credentials.Email)

	if err != nil || user == nil || ComparePassword(user.Password, credentials.Password) == false {
		http.Error(w, "Incorrect login credentials", http.StatusUnauthorized)
		return
	}

	RespondWithJSON(w, 200, struct{}{})
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	type SignupBody struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
	}

	var newUser SignupBody

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	existingUser, err := FindUserByEmail(newUser.Email)

	if err != nil || existingUser != nil {
		http.Error(w, "User with email already exist", http.StatusBadRequest)
		return
	}

	CreateUser(newUser.Email, newUser.Password)

	RespondWithJSON(w, 200, struct{}{})
}

func CreateLinks(w http.ResponseWriter, r *http.Request) {

	var newLinkBody Link

	err := json.NewDecoder(r.Body).Decode(&newLinkBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// link := Link{
	// 	URL:      newLinkBody.URL,
	// 	Platform: newLinkBody.Platform,
	// }
	// links = append(links, link)

	RespondWithJSON(w, 200, struct{}{})

}

func ListLinks(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, []string{})
}

func UpdateLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create links"))
}

func ReadProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create links"))
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create links"))
}
