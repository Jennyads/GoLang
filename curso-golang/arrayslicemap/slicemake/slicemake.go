package main

import "fmt"

func main() {
	s := make([]int, 10)  //slice
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	//com o make já se cria internamente um array associado 
	// não precisa colocar := pq ela essa variável já extiste, não está sendo introduzida
	// a variável está sendo reatrivuida com a função make que vai ter 10 elementos nesse slice, com um array interno/capacidade interna de 20 posições (elementos)
	// vai referenciar apenas 10
	fmt.Println(s, len(s), cap(s))
	// len tamanho do slice, cap para pegar capacidade maxima do slice (array interno)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	//quando adiciona novos elementos no slice, ele internamente vai referenciando array maiores e replicando os dados para o array maior
	fmt.Println(s, len(s), cap(s))
}