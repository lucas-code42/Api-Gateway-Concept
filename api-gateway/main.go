package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fmt.Println("START API-GATEWAY")

	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(callApi()))
	})

	http.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(callApi()))
	})

	http.ListenAndServe(":8080", nil)
}

func callApi() string {
	url := "http://127.0.0.1:8081/api"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	
	return string(body)
}
