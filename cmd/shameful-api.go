package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting Restful service")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8093", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RestfulServ. on:8093")
}
