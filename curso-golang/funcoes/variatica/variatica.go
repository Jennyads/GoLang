package main

import "fmt"
//três pontinhos representa função variatica que pode receber parâmetros variaveis
func media(numeros ...float64) float64 {
	total := 0.0 //array
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}