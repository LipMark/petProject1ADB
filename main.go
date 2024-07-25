package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"

	"github.com/LipMark/ToDo-Go/util"
)

var rnd *renderer.Render

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	filePath := "./README.md"
	err := rnd.FileView(rw, http.StatusOK, filePath, "readme.md")
	util.CheckError(err)
}

func main() {

	/*
		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Get("/", homeHandler)
		router.Mount("/todo", routes.TodoHandlers())
	*/
	server := &http.Server{
		Addr:         ":9000",
		Handler:      chi.NewRouter(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("Server started on port", 9000)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("listen:%s\n", err)
	}
}
