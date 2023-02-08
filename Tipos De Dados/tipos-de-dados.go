package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int8 = 100
	fmt.Println(numero)

	numero2 := 1000
	fmt.Println(numero2)

	var numero3 uint32 = 10000 // uint nao aceita numeros negativos
	fmt.Println(numero3)

	//alias
	//int 32 = RUNE

	var numero4 rune = 123456
	fmt.Println(numero4)

	//BYTE = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//FIM NUMERO REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := "b"
	fmt.Println(char)

	//FIM STRINGS

	var texto int16
	fmt.Println(texto)

	texto2 := 5
	fmt.Println(texto2)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error //ERRO NO GO
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno") //imprimir uma string de erro
	fmt.Println(erro2)                           // precisa de pacote ERRORS

}
