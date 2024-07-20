// не работает
// переход на другой проект

package main

import (
	"fmt"
	"net/http"
)

var (
	USERNAME = "12345"
	PASSWORD = "54321"
)

// "go-authentication-boilerplate/database"
//	"go-authentication-boilerplate/router"
// doesn't work too
//db
// postgre don't work?

// auth
func handleLogin(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something went wrong"))
		return
	}

	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		w.Write([]byte("wrong username/password"))
		return
	}

	w.Write([]byte("User logged in"))
}

//handlers func

func main() {
	//mux serv

	// serve-up
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
	// handlers call

	// mock
	fmt.Println("test project")

}

// new handlers again
