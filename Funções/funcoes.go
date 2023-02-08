package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
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
}
