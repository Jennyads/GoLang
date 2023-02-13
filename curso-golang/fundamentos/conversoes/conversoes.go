package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) // precisa converter o int para float 64

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println((notaFinal))

	//cuidado
	fmt.Println("Teste " + string(97)) //97 é o código corresponder da tabela ascii (unicode)
	//assim não converte int para string

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("teste", int(123))

	//string para int
	num, _ := strconv.Atoi(("123"))
	//num é o valor (STRING que será convertida), _ representa o erro ignorado (caso exista)
	fmt.Println((num - 122))
	// nesse caso converteu o NUM para 123

	//string para bolean
	b, _ := strconv.ParseBool("true") //pode substituir o true por 1
	if b {
		fmt.Println("Verdadeiro")
	}
}
