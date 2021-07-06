package main

import (
	"log"
	"net"

	wasp "../.."
)

func main() {
	server := wasp.Server{}
	err := server.Listen("localhost", 3333, func(message []byte, conn net.Conn) {
		// Print the incoming message
		log.Println(string(message))

		// Send a reply
		err := server.Send(conn, []byte("Hello from the server"))

		if err != nil {
			log.Fatal(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}
}
