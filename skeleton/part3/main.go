// Solution to part 3 of the Whispering Gophers code lab.
//
// This program listens on the host and port specified by the -listen flag.
// For each incoming connection, it launches a goroutine that reads and decodes
// JSON-encoded messages from the connection and prints them to standard
// output.
//
// You can test this program by running it in one terminal:
// 	$ part3 -listen=localhost:8000
// And running part2 in another terminal:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
)

var listenAddr = flag.String("listen", "localhost:8000", "host:port to listen on")

type Message struct {
	Body string
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("failed to listen to \"%s\" with tcp protocol: %v\n", *listenAddr, err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v\n", err)
			continue
		}

		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	d := json.NewDecoder(c)
	for {
		msg := Message{}
		if err := d.Decode(&msg); err == io.EOF {
			break
		} else if err != nil {
			log.Printf("failed to decode message: %v\n", err)
			continue
		} else {
			log.Printf("%v\n", msg)
		}
	}
}
