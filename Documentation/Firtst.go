package main

import (
	"example.com/hello/Documentation/Greetings" // example.com/hello (root) - path.
	"fmt"                                       // Pacote interno do go de formatação de textos.
)

func main() {

	message := Greetings.Hello("Tiagão")
	fmt.Println(message)
}
