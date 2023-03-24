package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail" //dependencia externa
)

func main(){

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	//auxiliar.escrever2() //aqui não pode ser exportada por conta de ser pacote diferente(letra minúscula)
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}

