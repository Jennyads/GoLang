package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}

//ter uma estrutura que recebe dados e le, forma de comunicação
//ponto de sincronismo