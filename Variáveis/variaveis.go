package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1" // declarando uma variavel declaando o tipo

	variavel2 := "Variavel 2" // declarando com o tipo automatico  (inferencia de tipo (tipo da variavel baseado no valor dele ))

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// outra forma de usar variaveis
	var (
		variavel3 string = "lalalalala"
		variavel4 string = "lelelelelel"
	)
	fmt.Println(variavel3, variavel4)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5 ", "variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	//no go voce pode inverter o valor das varivaeis sem problema nenhum sem precisar de uma outra variavel
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
