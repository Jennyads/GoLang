package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}

//closer seria fechamento, algo que envolve alguma coisa
//função que envolve outra função (função interna e externa)
//entende o local fisíco/contexto léxico e escopo que foi declarado
//a função externa funciona como closure (fechamento), vai ter a memória da origem

