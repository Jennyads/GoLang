package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2 //E, os dois tem que ocorrer de forma correta
	comprarTv32 := trab1 != trab2 // OU EXCLUSIVO, mas usa o diferente para representar pq não tem no go
	comprarSorvete := trab1 || trab2 //ou, qualquer um dos dois

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}
