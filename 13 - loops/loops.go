package main

import (
	"fmt"
)

func main() {

	//numero := 2

	/*for numero < 10 {
		numero++
		fmt.Println("teste", numero)
		time.Sleep(time.Second)
	}*/

	nomes := [3]string{"teste 1", "teste 2", "teste 3"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

}
