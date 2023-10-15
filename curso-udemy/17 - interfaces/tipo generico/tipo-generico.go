package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generica("teste")
	generica(123)

	mapa := map[int32]interface{}{
		1:   "teste",
		123: "tete",
	}

}
