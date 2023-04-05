package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)  //canal de comunicação, pode enviar ou receber dados
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")
	//<- canal, assim estaria recebendo um valor
	// mensagem := <-canal
	// fmt.Println(mensagem)

	// for{
	// 	mensagem, aberto := <-canal
	// 	if !aberto{
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
	
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto  //seta dessa forma significa que estou mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal) //evita dar deadlock fechando o canal, pois fica esperando dado sem ter, então é necessário ter uma finalização
}
