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
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login now"))
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register"))
}

func CreateLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create links"))
}

func ListLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create links"))
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
