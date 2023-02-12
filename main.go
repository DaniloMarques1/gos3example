package main

import (
	"log"
	"net/http"

	"github.com/danilomarques1/gos3example/api/controller"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	userController, err := controller.NewUserController()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/user", applicationJson(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Save(w, r)
		case http.MethodGet:
			userController.FindByEmail(w, r)
		}
	}))

	log.Printf("server running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func applicationJson(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
}
