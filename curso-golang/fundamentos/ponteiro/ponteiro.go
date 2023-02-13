package main

import "fmt"

func main() {
	i := 1 //int8

	//Go não tem aritmética de ponteiros

	//ponteiro = referência de memória

	var p *int = nil
	//p aponta para o endereço de memoria do tipo int
	//nil é um nulo, não aponta para lugar nenhum
	p = &i // pegando o endereço da variável i e atribui o endereço ao ponteiro p
	*p++   // desreferenciando (pegando o valor), com o * acessa o valor (modificando), se fosse só p pegaria o endereço
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
