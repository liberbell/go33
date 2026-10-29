package crud_app

type Handler struct {
	db *DB
}

func NewHandler(db *DB) *Handler {
	return &Handler{db: db}
}
