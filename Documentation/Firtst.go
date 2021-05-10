package main

import (
	"example.com/hello/Documentation/Errors"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Errors: ")
	log.SetFlags(0)

	message, err := Errors.ErrHello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
