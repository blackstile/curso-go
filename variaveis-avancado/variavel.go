package main

import "fmt"

//Endereco é uma variavel exportada (publica), pois ela comeca com letra maiuscula.
var Endereco string

//Telefone é uma variavel publica também
var Telefone string //definindo mais de uma variavel na mesma linha

///definindo variaveis em blocos
var (
	quantidade int
	comprou    bool
	valor      float32
	palavras   rune
)

func main() {
	//shorthand deve ser utilizados apenas dentro de funcao
	teste := "Valor de variavel teste" //shorthand define e atribui o valor, sem necessidade da palavra chave var

	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)

	fmt.Printf("teste: %s\r\n", teste)

}
