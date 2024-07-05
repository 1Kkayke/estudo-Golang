package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sorteioNumero()
}

func sorteioNumero() {
	var num1, num2 int

	num1 = rand.Int()

	fmt.Println("Digite um numero entre 0 e 2 para o sorteio")
	fmt.Scan(&num2)

	if num2 < 0 || num2 > 2 {
		fmt.Println("Numero invalido")
		return
	}

	if num1 == num2 {
		fmt.Println("Parabens, voce acertou")
	} else {
		fmt.Println("Que pena, voce errou, continue tentando")
		sorteioNumero()
	}
}
