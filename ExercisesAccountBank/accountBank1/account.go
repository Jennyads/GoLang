package main

import "fmt"

type balance struct{
	saldo           float64
	amountDebited   float64
	amountDeposited float64
	
}


func (b balance) getBalance() float64 {
	return b.saldo
}

func (b * balance) depositValue() {
	b.saldo += b.amountDeposited
	
}

func (b * balance) debitValue() {
	if b.saldo < b.amountDebited {
		fmt.Println("Saldo Insuficiente")
	} else {
		b.saldo -= b.amountDebited
		
	}
	
}

func main () {
	b1 := balance{saldo: 100, amountDeposited: 300, amountDebited: 100}
	b1.depositValue()
	b1.debitValue()
	fmt.Println(b1.getBalance())

}

