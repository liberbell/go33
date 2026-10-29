package crud_app

import "net/http"

type Handler struct {
	db *DB
}

func NewHandler(db *DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
}
