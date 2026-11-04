package main

import "todolist/todolist"

func main() {
	db := todolist.NewDatabse()
	db.InitDB()
	defer db.CloseDB()

	handlers := todolist.Handler{DB: db}
	router := handler.InitRouter()
}
