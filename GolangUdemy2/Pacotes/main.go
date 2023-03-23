package main

import (
	"modulo/auxiliar"
	"fmt")

func main(){

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	//auxiliar.escrever2() //aqui não pode ser exportada por conta de ser pacote diferente(letra minúscula)
}

