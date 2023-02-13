package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice == tem tamanho va
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slide define um pedaço de um array.
	s2 := a2[1:3] //não inclui o indice 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	//vai do inicio até o indice 2 sem incluir o mesmo
	fmt.Println(a2, s3)

	// vc pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] //slice de um slice
	//inicio até o indice 1 sem incluir o mesmo
	fmt.Println(s2, s4)
}