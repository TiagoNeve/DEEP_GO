package Errors

import (
	"errors"
	"fmt"
)

/*ErrHello Função recebe uma string nome e
  retorna uma saudação.
*/
func ErrHello(nome string) (string, error) {
	// Se o nome estiver vazio retorna a string vazio e mostra um error
	if nome == "" {
		return "", errors.New("nome vazio, digite um nome válido")
	}
	message := fmt.Sprintf("Eae, %v. Bem vindo!", nome)
	return message, nil
}
