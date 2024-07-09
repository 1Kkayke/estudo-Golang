package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	initialize()
}

func initialize() {
	server := gin.Default()
	server.GET("/calculator", getParams)
	server.Run()
}

func getParams(c *gin.Context) {
	number1 := c.Query("number1")
	number2 := c.Query("number2")
	typeOperation := c.Query("operation")

	if number1 == "" || number2 == "" || typeOperation == "" {
		c.String(http.StatusBadRequest, jsonError("Preencha todos os campos"))
		if strings.ToUpper(typeOperation) != "SOMA" && typeOperation != "SUBTRAI" && typeOperation != "MULTIPLICA" && typeOperation != "DIVISAO" {
			c.String(http.StatusBadRequest, jsonError("Operação inválida"))
		}
	} else {
		c.String(http.StatusOK, calculate(number1, number2, typeOperation))
	}

}

func calculate(number1 string, number2 string, typeOperation string) string {

	number1Int, _ := strconv.Atoi(number1)
	number2Int, _ := strconv.Atoi(number2)
	result := 0

	if strings.ToUpper(typeOperation) == "SOMA" {
		result = number1Int + number2Int
	} else if typeOperation == "SUBTRAI" {
		result = number1Int - number2Int
	} else if typeOperation == "MULTIPLICA" {
		result = number1Int * number2Int
	} else if typeOperation == "DIVISAO" {
		result = number1Int / number2Int
	}

	return strconv.Itoa(result)
}

func jsonError(message string) string {
	return "{\"error\":\"" + message + "\"}"
}
