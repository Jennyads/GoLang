package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta e gera a quantidade correta de elementos!
	//[,,,] é um array com número fixo, porém só sera contado com base/a partir da quantidade de elementos que for colocada na inicialização do array
	// se tirar o [...] não será array, mas sim slice

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
	//_ ignora o valor dos índices
		fmt.Println(num)
	}
}
