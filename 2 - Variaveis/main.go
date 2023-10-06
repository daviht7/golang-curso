package main

import "fmt"

func main() {

	var variavel1 string = "variavel 1 "

	variavel2 := "variavel2"

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	variavel5, variavel6 := "teste", "teste"

	const consten string = "constante"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(consten)

}
