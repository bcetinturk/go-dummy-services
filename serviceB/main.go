package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
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
	Port: getenv("SERVICEB_PORT", "8081"),
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Got /message Request")

	messages := []string{"Hello, how are you?", "Hallo, wie geht's dir?", "¿Hola, cómo estás?", "Hallo hoe is het?"}
	index := rand.Intn(len(messages))

	message := messages[index]
	log.Printf("Picked: %s\n", message)
	io.WriteString(w, message)
}

func main() {
	http.HandleFunc("/message", getMessage)

	err := http.ListenAndServe(fmt.Sprintf(":%s", appConfig.Port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
