package main

import "fmt"

func main() {
	calculoNota()
}
func calculoNota() {

	var nota1, nota2, media float64
	var nome string
	fmt.Println("Qual o seu nome?")
	fmt.Scan(&nome)

	fmt.Println("Qual sua primeira nota?")
	fmt.Scan(&nota1)
	if nota1 < 0 || nota1 > 10 {
		fmt.Println("Nota invalida")
		return
	}

	fmt.Println("Qual sua segunda nota?")
	fmt.Scan(&nota2)
	if nota2 < 0 || nota2 > 10 {
		fmt.Println("Nota invalida")
		return
	}

	media = (nota1 + nota2) / 2
	fmt.Printf("A media de %s foi %.2f\n", nome, media)
}
