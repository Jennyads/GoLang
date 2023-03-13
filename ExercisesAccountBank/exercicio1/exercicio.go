//Faça um programa que sirva para calcular o lucro liquido de uma loja, leve em consideração as seguintes regras.
//Imposto 30%, Funcionário 10% sobre o resultado do imposto. 3% de ICMS se for um mês par e 2% se for impar.

package main 

import "fmt"


func imposto(lucroBruto float64) float64 {
	return lucroBruto * 0.7

}
func taxaFuncionario(valorImposto float64) float64{
	return valorImposto * 0.9

}

func main() {

	var lucroBruto float64

	fmt.Println("Informe o lucro bruto da loja: ")
	fmt.Scanln(&lucroBruto)



	valorImposto := imposto(lucroBruto)
	valorFuncionario := taxaFuncionario(valorImposto)

	var mes int
	fmt.Println("Informe o mês atual (1 a 12): ")
	fmt.Scanln(&mes)
	var lucroLiquido float64
	if mes%2 == 0 {
		lucroLiquido = valorFuncionario * 0.97
	} else {
		lucroLiquido = valorFuncionario * 0.98
	}

	fmt.Printf("O lucro líquido é de R$ %.2f", lucroLiquido)


}