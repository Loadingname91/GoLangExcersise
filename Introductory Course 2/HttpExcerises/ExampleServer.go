package main

import (
	"fmt"
	"net/http"
)

func Examplemain() {

	mux := http.NewServeMux()

	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recieved")
		w.Write([]byte("Hello world"))
		fmt.Println(r.Method)
	})

	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recieved")
		w.Write([]byte("Hello world go"))
		fmt.Println(r.Method)
	})

	fmt.Println("listiing started")
	http.ListenAndServe("localhost:8080", mux)
}
