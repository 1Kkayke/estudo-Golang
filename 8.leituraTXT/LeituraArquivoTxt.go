package main

import (
	"fmt"
	"log"
	"os"

	// "log"

	"github.com/sqweek/dialog"
)

func main() {
	getArchive()
}

func getArchive() {

	var (
		filePath string
		getFile  *os.File
	)
	filePath, err := dialog.File().Title("Selecione um arquivo").Load()

	if err != nil {
		log.Fatal(err)
	}

	getFile, err = os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer getFile.Close()

	fmt.Println("Arquivo selecionado:", filePath)

	ReadArchive(getFile)
}

func ReadArchive(getFile *os.File) {
	var (
		leitor []byte
		err    error
	)

	leitor, err = os.ReadFile(getFile.Name())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(leitor))

	getFile.Close()
}
