package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application-json")
	type Greetings struct {
		Data string `json:"data"`
	}
	message := "hello world"
	response, err := json.Marshal(Greetings{Data: message})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(response)

}

func main() {
	fmt.Println("Hello, welcome to links backend")

	router := chi.NewRouter()
	router.Get("/health", Health)
	router.Post("/login", HandleLogin)
	router.Post("/register", HandleRegister)
	router.Post("/create-links", CreateLinks)
	router.Get("/list-links", ListLinks)
	router.Get("/profile", ReadProfile)
	router.Patch("/profile/update", UpdateProfile)

	server := http.Server{
		Handler: router,
		Addr:    ":8081",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
