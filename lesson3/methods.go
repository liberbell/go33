package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello" + u.Name
}

func (u User) UpperName() {
	u.Name = strings.ToUpper(u.Name)
}

func (u *User) setName(name string) {
	u.Name = name
}

func main() {
	user := User{Name: "James"}
	fmt.Println(user.Greet())
}
