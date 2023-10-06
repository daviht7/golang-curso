package main

import "fmt"

func main() {
	fmt.Println("teste")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	variavel3++

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	fmt.Println(variavel3)
	fmt.Println(*ponteiro)

}
