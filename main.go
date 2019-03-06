package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	p := Person{"petesh", 30}
	encoded, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	p1 := &Person{}
	err = json.Unmarshal(encoded, p1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v",p1)
}