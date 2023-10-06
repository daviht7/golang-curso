package main

import (
	"crypto"
	"errors"
	"fmt"
)

func main() {

	numero := 1000000
	var negativo uint = 123
	var rune rune = 123 //int32
	var byte byte = 123
	var teste3 float32 = 123

	var variavelnaodeclarada float32

	char := 'b'

	var isTeste bool = teste3 == 123

	var erro error = errors.New("teste3 must be zero or more")

	t := crypto.BLAKE2b_256.New()

	//float32 float64

	fmt.Println(numero)
	fmt.Println(negativo)
	fmt.Println(rune)
	fmt.Println(byte)
	fmt.Println(teste3)
	fmt.Println(char)
	fmt.Println(variavelnaodeclarada)
	fmt.Println(isTeste)
	fmt.Println(erro)
	fmt.Println(t)

}
