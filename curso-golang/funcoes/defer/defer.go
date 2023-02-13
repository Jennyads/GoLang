package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	//vai ser atrasada até o momento antes do retorno da função
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} 
	fmt.Println("Valor baixo...") //assim é uma alternativa para o else
	return numero
	
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
