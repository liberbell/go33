package main

import (
	"crud_app/crud_app"
	"log"

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
}
