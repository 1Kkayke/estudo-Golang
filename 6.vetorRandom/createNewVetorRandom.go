package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	CriaNovoVetor()
}

func CriaNovoVetor() {

	vector := rand.NewSource(time.Now().UnixNano())

	x := rand.New(vector)

	for i := 0; i < 11; i++ {
		// Generate a random number
		numerosAleatorios := x.Intn(100)
		fmt.Println(numerosAleatorios)
	}
}
