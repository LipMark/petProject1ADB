// не работает
// переход на другой проект

package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// "go-authentication-boilerplate/database"
//	"go-authentication-boilerplate/router"
// doesn't work too
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

// templ from git
func home(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		toDos := todolist.LoadAllToDo()

		err := templates.ExecuteTemplate(rw, "home", types.LoadAllToDoData{ToDos: toDos})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
		break
	case "POST":
		res, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		id := fmt.Sprintf("%s", res)
		getID := strings.Split(id, "=")
		toDoID, err := strconv.Atoi(getID[1])
		if err != nil {
			panic(err.Error())
		}
		todolist.DoneToDo(toDoID)
		http.Redirect(rw, r, "/", http.StatusFound)
		break
	default:
		break
	}
}
