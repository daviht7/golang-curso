package main

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	var u estudante

	u.altura = 10

}
