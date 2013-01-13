package main

import (
	"time"
	"fmt"
	"net/http"
	"github.com/gorilla/schema"
)

type User struct {
	Id int `PK`
	Email string `schema:"email"`
	Password string `schema:"password"`
	Status int `schema:"-"`
	Last_login time.Time `schema:"-"`
	Created time.Time `schema:"-"`
}

func NewUser () *User {
	return &User{}
}

func (u *User) Init () {
	http.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		_ = u.Register(r)
	})

	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
		u.Email = "test@test.com"
		u.Save()
	})
}

func (u *User) Save () {
	game.db.orm.Save(u)
}

func (u *User) Register (r *http.Request) int {
	decoder := schema.NewDecoder();
	decoder.Decode(u, r.Form)
	fmt.Print(u);
	return 0
}
