package prefixo

import (
	"strconv"

	"github.com/blackstile/curso-go/pacotes/operadora"
)

//Capital prefixo da capital
var Capital = 11

//NomeOperadoraPrefixoInterior  nome da operadora mais o código prefixo interior
var NomeOperadoraPrefixoInterior = operadora.NomeDaOperadora + " - 13"

//variavel privada, pode ser utilizada apenas dentro deste arquivo
var teste = "teste"

//TesteComPrefixo é apenas um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
