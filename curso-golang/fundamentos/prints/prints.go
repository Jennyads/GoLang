package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	//fmt.Print(("O valor de x é " + x))
	//não permite concatenação de float com string
	xs := fmt.Sprint(x)
	fmt.Println(("O valor de x é " + xs)) //alternativa tendo que criar a váriavel
	
	fmt.Println("O valor de x é", x)      //alternativa direta

	fmt.Printf("O valor de x é: %.2f.", x) //print formatado, com 2 casas decimais
	// %f imprimir float, %s imprimir string, %d inteiro, %t bolean, %v é mais genérico

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
