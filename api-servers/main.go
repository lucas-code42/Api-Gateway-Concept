package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("START API SERVER")

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api"))
	})

	http.ListenAndServe(":8081", nil)
}
