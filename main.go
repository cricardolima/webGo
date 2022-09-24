package main

import (
	"net/http"
	"webGo/main.go/routes"
)

func main() {
	routes.LoadRoutes()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}
