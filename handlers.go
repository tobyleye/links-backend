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
	"fmt"
	"net/http"
)

type User struct {
	ID       string
	Email    string
	Password string
}

type Link struct {
	Platform string
	URL      string
}

type Users []User

type Links []Link

var links Links = []Link{}
var users Users = []User{}

func findUserByEmail(email string) User {
	var searchedUser User

	for _, user := range users {
		if user.Email == email {
			searchedUser = user
		}
	}
	return searchedUser
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
	var searchedUser = findUserByEmail(credentials.Email)

	if searchedUser.ID == "" || searchedUser.Password != credentials.Password {
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

	var existingUser = findUserByEmail(newUser.Email)
	if existingUser.ID != "" {
		http.Error(w, "User with email already exist", http.StatusBadRequest)
		return
	}
	nextUserId := len(users) + 1
	user := User{ID: fmt.Sprint(nextUserId), Email: newUser.Email, Password: newUser.Password}
	users = append(users, user)
	RespondWithJSON(w, 200, user)
}

func CreateLinks(w http.ResponseWriter, r *http.Request) {

	var newLinkBody Link

	err := json.NewDecoder(r.Body).Decode(&newLinkBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	link := Link{
		URL:      newLinkBody.URL,
		Platform: newLinkBody.Platform,
	}
	links = append(links, link)

	RespondWithJSON(w, 200, link)

}

func ListLinks(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, links)
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
