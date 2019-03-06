package main

import (
	"github.com/apetesh/gothings/greeter"
	"log"
)

func main() {
	srv := &greeter.Server{}
	log.Fatal(srv.Start("8000"))
}
