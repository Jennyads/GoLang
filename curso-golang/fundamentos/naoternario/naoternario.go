package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado" assim não existe
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}