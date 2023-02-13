
package main

import "runtime/debug" 


func f3() {
	debug.PrintStack()
	//vai imprimir a pilha de execução de um programa no momento que essa função for chamada

}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
	//main, f1, f2, f3 
	//pilha que coloca e tira metodos
}