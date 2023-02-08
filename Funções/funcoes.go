package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
func calculosMatematicos(n1,n2 int8) (int8, int8){     // permitido ter mais de um retorno por funçao
	soma := n1 +n2                                      //muito utilizado em tratamento de erro
	subtracao := n1 -n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Funçao f") // exemplo 1
	}
	var f2 = func(txt string) { //exemplo 2
		fmt.Println(txt)
	}

	f()
	f2("Texto da funcao ")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10,15)
	fmt.Println(resultadoSoma,resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(10,15)       // calcula 2 mais so mostra  um 
	fmt.Println(resultadoSoma2,)
}
