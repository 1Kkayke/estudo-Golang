package main

import "fmt"

func main() {

	calculoSalarioFuncionario()
}

func calculoSalarioFuncionario() {

	var nome string
	var horasTrabalhadas, valorPorHora float64
	var salario float64

	fmt.Println("Qual o nome do funcionario?")
	fmt.Scan(&nome)

	fmt.Println("Quantas horas o funcionario trabalhou?")
	fmt.Scan(&horasTrabalhadas)

	fmt.Println("Qual o valor por hora trabalhada?")
	fmt.Scan(&valorPorHora)

	salario = horasTrabalhadas * valorPorHora
	fmt.Printf("O salario de %s e de %.2f\n", nome, salario)

}
