package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello, World!"
		w.Write([]byte(message))
	})

	address := ":8080"

	fmt.Println("Server is running on", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
