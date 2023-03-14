package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	ServiceBURL string
	Port        string
}

func getenv(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}

var appConfig = Config{
	ServiceBURL: getenv("SERVICEA_SERVICEB_URL", "127.0.0.1:8081"),
	Port:        getenv("SERVICEA_PORT", "8080"),
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Got / Request")
	io.WriteString(w, "Check out /hello !\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Got /hello Request")

	log.Printf("Trying %s\n", appConfig.ServiceBURL)
	resp, _ := http.Get(fmt.Sprintf("http://%s/message", appConfig.ServiceBURL))

	body, _ := io.ReadAll(resp.Body)

	sb := string(body)
	log.Printf("Got %s from serviceB\n", sb)
	io.WriteString(w, fmt.Sprintf("%s\n", sb))
}

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(fmt.Sprintf(":%s", appConfig.Port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
