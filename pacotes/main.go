package main

import (
	"fmt"

	"github.com/blackstile/curso-go/pacotes/operadora"
	"github.com/blackstile/curso-go/pacotes/prefixo"
)

//NomeDoUsuario define o nome do usuario do sistema
var NomeDoUsuario = "William"

func main() {
	fmt.Printf("Nome do Usuário: %s \r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeDaOperadora)
	fmt.Printf("Nome da operadora: %s\r\n", prefixo.TesteComPrefixo)
}
