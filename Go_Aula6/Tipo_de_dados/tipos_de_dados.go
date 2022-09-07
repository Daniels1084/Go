package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int64 = -10000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//int32 = RUNE

	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8

	var numero4 byte = 123
	fmt.Println(numero4)

	//Fim dos numeros inteiros
	//Número Reias

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123333300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//Fim numeros reais
	//String

	var texto string = "Texto"

	fmt.Println(texto)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// Fim Strings

	texto1 := 5
	fmt.Println(texto1)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}
