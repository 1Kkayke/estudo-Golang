package main

import "fmt"

func reverse(stringUser string) (resultado string) {
	for _, v := range stringUser {
		resultado = string(v) + resultado
	}
	return
}

func main() {

	stringUser := ""

	fmt.Scanln(&stringUser)

	stringRevertida := reverse(stringUser)
	fmt.Println(stringUser)
	fmt.Println(stringRevertida)
}
