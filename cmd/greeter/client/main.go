package main

import (
	"github.com/apetesh/gothings/greeter"
	"log"
)

func main() {
	client := greeter.NewClient("http://127.0.0.1:8000")
	msg, err := client.Greet("petesh")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response: %s", msg.Message)
}
