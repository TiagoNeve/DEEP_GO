package Greetings

import "fmt"

/*Hello função recebe um nome em string e retorna uma mensagem de saudação

	Exemplo: mensagem := Hello("Tiago")
             fmt.Println(mensagem)

             output: Olá, Tiago. Bem vindo!
*/
func Hello(name string) string {
	message := fmt.Sprintf("Olá, %v. Bem vindo(a)!", name)
	return message
}
