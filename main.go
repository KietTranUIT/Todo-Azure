package main

import (
	"log"
	"net/http"

	"github.com/ichtrojan/go-todo/routes"
)

func main() {
	err := http.ListenAndServe(":8080", routes.Init())

	if err != nil {
		log.Println("Loi: ", err.Error())
	}
}
