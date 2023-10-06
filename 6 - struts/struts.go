package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Print("teste")

	var u usuario
	u.nome = "teste"
	u.idade = 123
	fmt.Println(u)

	u2 := usuario{"teste", 123}

	fmt.Println(u2)

}
