package main

import "fmt"          // Pacote interno do go de formatação de textos.
import "rsc.io/quote" // Importando pacotes diretamente da internet.

func main() {
	//OlaMundo()
	//QuoteCall()
	testFmt()
}

func testFmt() {
	fmt.Println("Aqui mostra um texto normal")
}

func OlaMundo() {
	fmt.Println("Eae mundão!!!")
}

func QuoteCall() {
	fmt.Println("Go:", quote.Go())         // Funciona perfeitamente e adiciona um espaço automático.
	fmt.Println("Glass: " + quote.Glass()) // Apenas concatena.
	fmt.Println("Hello: " + quote.Hello())
	fmt.Println("Opt:", quote.Opt()) // Que frase em parça
}
