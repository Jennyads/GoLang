package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	// float 64 tipo padrão de um literal com ponto flutuante(numero real), foi inferido pelo compilador
	var raio = 3.12
	// em gol não é obrigado a colocar o tipo da variável se atribui um valor a ela
	//o compilador consegue fazer a inferência do valor
	area := PI * math.Pow(raio, 2) // raio ao quadrado == math.Pow é usado para cálculo de potência
	// := forma reduzida,é como se a váriavel ainda não existisse, mas eu estou atribuindo um valor e a inicializando.
	//assim declara e inicializa numa construção só, de maneira reduzida (forma mais usada)
	fmt.Println("A área da circunferência é", area)


    // bloco de construção de constantes
	const (
		a = 1
		b = 2
	)
	// bloco de construção de variáveis
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
