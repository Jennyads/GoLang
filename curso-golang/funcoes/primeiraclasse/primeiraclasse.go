package main

import "fmt"


// por que nessa função eu não posso inicializar com :=?

//soma := func(a, b int) int  => não entendi pq não pode ser declarada e inicializada dessa forma
//por que só dentro do main que pode?


var soma = func(a, b int) int { //função anônima, sem nome
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
