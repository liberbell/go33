package main

import (
	"crud_app/crud_app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := crud_app.NewDatabases()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	handler := crud_app.NewHandler(db)
	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")

	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
