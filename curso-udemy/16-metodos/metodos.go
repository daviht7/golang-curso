package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u *usuario) salvar() {
	//fmt.Print("salvando...")
	u.idade = 15
}

func escrever() {

	fmt.Println("escrevendo..")

	usuario1 := usuario{"davi", 8}

	fmt.Println(usuario1)

	usuario1.salvar()

	fmt.Println(usuario1)

}

func main() {
	escrever()
}
