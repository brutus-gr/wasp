package main

import (
	"log"

	wasp "../.."
)

func main() {
	client := wasp.Client{}
	err := client.ConnectTo("localhost", 3333)

	if err != nil {
		log.Fatal(err)
	}

	go client.Listen(func(message []byte) {
		// Print incoming messages
		log.Println(string(message))
	})

	// Send an outgoing message
	err = client.Send([]byte("Hello from the client"))

	if err != nil {
		log.Fatal(err)
	}

	for {
	}
}
