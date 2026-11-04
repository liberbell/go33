package main

import "todolist/todolist"

func main() {
	db := todolist.NewDatabse()
	db.InitDB()
	defer db.CloseDB()
}
