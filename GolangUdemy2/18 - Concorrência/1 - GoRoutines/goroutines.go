package main

import (
	"fmt"
	"time"
)
//concorrência é diferente de paralelismo 
//paralelismo  acontece quando duas ou mais tarfas são executadas ao mesmo tempo, isso só é possívem em processador com mais de um núcleo
//concorrente não necessariamente ao mesmo tempo, uma tarefa não espera a outra acabar, elas se revezam durante o tempo
func main() {
	go escrever("Olá Mundo!") // goroutine - metodo que é envocado para não esperar finalizar a função e ir para a próxima, assim alterna
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
