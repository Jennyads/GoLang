package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") //vai tentar excluir, mas nao vai acontecer nada

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	//ingorar o nome
	for _, salario := range funcsESalarios {
		fmt.Println(salario)
	}

	//vai imprimir só o nome que é o primeiro valor que será a chave representada
	for salario := range funcsESalarios {
		fmt.Println(salario)
	}
}