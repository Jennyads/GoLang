package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

//essa função vai receber como primeiro parâmetro uma função tendo dois parâmetros (assinatura) inteiro e retonra um inteiro 

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}