package controller

import (
	CustomStructs "MyGoCourse/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := CustomStructs.MyResponse{
				Code: http.StatusOK,
				Body: "hellow rold",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}

		fmt.Println("request recieved")
		w.Write([]byte("Hello world"))
		fmt.Println(r.Method)
	}
}
