package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		// Allow any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow all methods
		w.Header().Set("Access-Control-Allow-Methods", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			// Handle preflight requests here
			w.WriteHeader(http.StatusOK)
			return
		}

		// Allow specified headers

		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Hello, welcome to links backend")

	router := chi.NewRouter()

	router.Use(enableCors)
	router.Get("/health", Health)
	router.Post("/login", HandleLogin)
	router.Post("/register", HandleRegister)
	router.Post("/links/create", CreateLinks)
	router.Get("/links/list", ListLinks)
	router.Get("/profile", ReadProfile)
	router.Patch("/profile/update", UpdateProfile)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	server := http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server is running on %s\n", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
