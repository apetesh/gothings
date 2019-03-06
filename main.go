package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Age int     `json:"age"`
	Name string `json:"name"`
	Parents struct {
		Mother &Person
		Father Person
	}
}

func main() {

}
