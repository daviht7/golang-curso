package main

import "fmt"

func main() {

	fmt.Println("teste")

	var array1 [5]int

	array1[0] = 10

	array2 := [5]string{"teste"}

	array3 := [...]int{123, 123, 123, 123}

	slice := []int

	fmt.Println(array1)

}
