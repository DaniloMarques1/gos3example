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

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Save(w, r)
		case http.MethodGet:
			// TODO
		}
	})

	log.Printf("server running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
