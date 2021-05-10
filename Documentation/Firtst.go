package main

import (
	"fmt"
)                     // Pacote interno do go de formatação de textos.
import "rsc.io/quote" // Importando pacotes diretamente da internet.

func main() {
	//OlaMundo()
	//QuoteCall()
	testFmt()
}

func testFmt() {
	// fmt.Println("Aqui mostra um texto normal")

	//fmt.Println(fmt.Errorf("essa é uma mensagem de error que pode ser formatada"))

	// fmt.Print("Escreve um texto normal que pode ter variaveis separadas por ,")
	// fmt.Print("E não pula a linha, padrão")

	// n, err := fmt.Fprint(os.Stdout, "Verifica a quantidade de bytes em um texto")

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
