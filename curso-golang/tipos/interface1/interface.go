package main

import "fmt"

type imprimivel interface { 
	toString() string //para a estrutura ser imprimivel ou não, ela precisa ter o metodo toString associado a ela
}
//estruturas que tem metodos que respeitam o metodo da interface, possiblitando o polimorfismo

//polimorfismo, múltiplas formas (pessoa e produto) a partir da interface imprimivel

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)
    //a mesma variavel coisa consegue ter multiplas formas, tanto como produto ou pessoa
	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)
}