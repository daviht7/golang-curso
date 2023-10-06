package main

import "fmt"

func main() {

	fmt.Print(diaDaSemana(10))

}

func diaDaSemana(numero int) string {

	switch numero {

	case 1:
		{
			return "segunda"
		}
	default:
		{
			return "terça"
		}

	}

}
