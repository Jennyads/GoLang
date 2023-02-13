package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros

	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000)) // indica o tipo

	//sem sinal (só positivos)... int8 int16 int32 int64

	//byte é um apelido para int8
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) //negação, vai passar a ser falso

	//string
	s1 := "Olá, meu nome é Jenny"
	fmt.Println(s1 + "!") //concatenação
	fmt.Println("O tamanho da string é", len(s1))

	//string com múltiplas linhas
	s2 := `Olá
	meu
	nome 
	é 
	Jenny`
	fmt.Println("O tamanho da string é", len(s2))

	//char, não tem um valor de char, pq na verdade é um valor do tipo int32
	//var x char = 'b' === NÃO EXISTE!
	char := 'a' //chama o a referente a tabela unicode
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}
