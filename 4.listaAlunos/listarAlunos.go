package main

import "fmt"

func main() {

	var alunos [5]string
	fmt.Println("Digite o nome dos alunos")
	for i := 0; i < 5; i++ {
		fmt.Scan(&alunos[i])
	}

	fmt.Println("Alunos: ", alunos)
}
