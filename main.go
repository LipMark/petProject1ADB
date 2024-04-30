package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// "go-authentication-boilerplate/database"
//	"go-authentication-boilerplate/router"

//db
// postgre don't work?

// auth

//handlers func

func main() {
	//listen serve chi?base?
	r := chi.NewRouter()
	// serve-up
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
	// handlers call

	// mock
	fmt.Println("test project")
}
