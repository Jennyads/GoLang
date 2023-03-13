package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Conta é a estrutura que representa uma conta bancária
type Conta struct {
	nome      string
	numero    int
	saldo     float64
	senha     int
}

// Contas é um slice que armazena todas as contas bancárias
var Contas []Conta

func main() {
	// Adicionar algumas contas de exemplo
	Contas = []Conta{
		{nome: "João Silva", numero: 1, saldo: 1000, senha: 123456},
		{nome: "Maria Oliveira", numero: 2, saldo: 2000, senha: 654321},
		{nome: "Pedro Souza", numero: 3, saldo: 3000, senha: 111111},
	}

	// Mostrar menu de opções
	fmt.Println("1. Consultar saldo")
	fmt.Println("2. Depositar")
	fmt.Println("3. Debitar")
	fmt.Println("4. Sair")
	fmt.Print("Escolha uma opção: ")

	// Ler a opção do usuário
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	opcao, _ := strconv.Atoi(text)

	// Executar a ação escolhida pelo usuário
	switch opcao {
	case 1:
		consultarSaldo()
	case 2:
		depositar()
	case 3:
		debitar()
	case 4:
		fmt.Println("Obrigado por usar o nosso sistema!")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida")
		main()
	}
}

func consultarSaldo() {
	// Solicitar número da conta e senha
	fmt.Print("Número da conta: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	numero, _ := strconv.Atoi(text)

	fmt.Print("Senha: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	senha, _ := strconv.Atoi(text)

	// Procurar a conta com o número fornecido
	var conta Conta
	encontrado := false
	for i := 0; i < len(Contas); i++ {
		if Contas[i].numero == numero && Contas[i].senha == senha {
			conta = Contas[i]
			encontrado = true
		break
		}
	}

	// Verificar se a conta foi encontrada
	if !encontrado {
		fmt.Println("Conta não encontrada ou senha incorreta")
		main()
		return
	}

	// Mostrar o saldo da conta
	fmt.Printf("O saldo da conta de %s é de R$ %.2f\n", conta.nome, conta.saldo)
	main()
}

func depositar() {
	// Solicitar número da conta e senha
	fmt.Print("Número da conta: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	numero, _ := strconv.Atoi(text)
	fmt.Print("Senha: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	senha, _ := strconv.Atoi(text)

	// Procurar a conta com o número fornecido
	var conta Conta
	encontrado := false
	for i := 0; i < len(Contas); i++ {
		if Contas[i].numero == numero && Contas[i].senha == senha {
			conta = Contas[i]
			encontrado = true
			break
		}
	}

	// Verificar se a conta foi encontrada
	if !encontrado {
		fmt.Println("Conta não encontrada ou senha incorreta")
		main()
		return
	}

	// Solicitar o valor a ser depositado
	fmt.Print("Digite o valor a ser depositado: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	valor, _ := strconv.ParseFloat(text, 64)

	// Depositar o valor na conta
	conta.saldo += valor

	// Mostrar o saldo atualizado da conta
	fmt.Printf("Depósito realizado com sucesso. O novo saldo é de R$ %.2f\n", conta.saldo)
	main()

}

func debitar() {
	// Solicitar número da conta, senha e valor a debitar
	fmt.Print("Número da conta: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	numero, _ := strconv.Atoi(text)
	fmt.Print("Senha: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	senha, _ := strconv.Atoi(text)

	fmt.Print("Valor a debitar: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	valor, _ := strconv.ParseFloat(text, 64)

	// Procurar a conta com o número fornecido
	var conta Conta
	encontrado := false
	for i := 0; i < len(Contas); i++ {
		if Contas[i].numero == numero && Contas[i].senha == senha {
			conta = Contas[i]
			encontrado = true
			break
		}
	}
	
	// Verificar se a conta foi encontrada
	if !encontrado {
		fmt.Println("Conta não encontrada ou senha incorreta")
		main()
		return
	}

	// Verificar se o saldo é suficiente
	if conta.saldo < valor {
		fmt.Println("Saldo insuficiente")
		main()
		return
	}
	
	// Debitar o valor da conta
	conta.saldo -= valor
	fmt.Printf("R$ %.2f foram debitados da conta de %s\n", valor, conta.nome)
	main()
	
	// Mostrar o saldo atualizado da conta
	fmt.Printf("Débito realizado com sucesso. O novo saldo é de R$ %.2f\n", conta.saldo)
	main()
}

