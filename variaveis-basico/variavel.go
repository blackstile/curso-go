package main

import (
	"fmt"
)

//variaveis definidas fora de uma função são variaveis de pacotes
//o padrao da definicao de variavel é o nome da variavel e depois o tipo
//o GO ja inicializa as variáveis por padrao para você

var endereco string // valor padrao de endereco= ""
var telefone = "9999-99999"
var quantidade int //quantidade=0
var comprou bool   //comprou= false
var valor float64  //valor = 0.0
var palavras rune

func main() {
	fmt.Printf("endereço: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)

}
