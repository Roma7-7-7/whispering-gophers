// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	encoder := json.NewEncoder(os.Stdout)

	for scanner.Scan() {
		msg := Message{
			Body: scanner.Text(),
		}
		if err := encoder.Encode(msg); err != nil {
			log.Fatalf("failed to encode message \"%s\": %v\n", msg, err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input: %v\n", err)
	}
}
