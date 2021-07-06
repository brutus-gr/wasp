# wasp

Wasp is a simple library for transmitting byte slices over TCP.

### Client usage
```golang
package main

import (
  "log"

  "github.com/brutus-gr/wasp"
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

```

### Server usage
```golang
package main

import (
  "log"
  "net"

  "github.com/brutus-gr/wasp"
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

```
