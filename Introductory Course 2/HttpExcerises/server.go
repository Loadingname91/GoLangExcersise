package main

import (
	"MyGoCourse/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:8080", mux)
}
