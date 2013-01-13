package main

import (
	"time"
	"fmt"
	"net/http"
)

type User struct {
	Id int `PK`
	Email string
	Password string
	Status int
	Last_login time.Time
	Created time.Time
}

func NewUser () *User {
	return &User{}
}

func (u *User) Init () {
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
		u.Email = "test@test.com"
		u.Save()
	});
}

func (u *User) Save () {
	game.db.orm.Save(u)
}
