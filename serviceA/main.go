package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Got / Request")
	io.WriteString(w, "Check out /hello !\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Got /hello Request")

	resp, _ := http.Get("http://127.0.0.1:8081/message")

	body, _ := io.ReadAll(resp.Body)

	sb := string(body)
	log.Printf("Got %s from serviceB\n", sb)
	io.WriteString(w, fmt.Sprintf("%s\n", sb))
}

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
