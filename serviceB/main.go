package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

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

	err := http.ListenAndServe(":8081", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
