package main

import "fmt"

func main(){
	var a int = 42
	var ptr *int = &a

	fmt.Printf("Endereço de a: %v\n", ptr)
	fmt.Printf("Valor de ptr: %v\n", a)
}