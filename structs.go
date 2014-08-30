package main

import (
	"fmt"
)

func main() {
	var carro Veiculo

	fmt.Println("Teste de structs")
	fmt.Println("Você possui um carro.")
	fmt.Print("Diga quantas portas ele tem: ")
	fmt.Scan(&carro.portas)
	fmt.Print("Diga quantas rodas ele tem: ")
	fmt.Scan(&carro.rodas)

	fmt.Println("Seu Carro tem ", carro.rodas, " rodas e ", carro.portas, " portas")

	carro.acelerar()
	carro.freiar()
	carro.re()
}

type Veiculo struct {
	rodas  int
	portas int
}

func (v Veiculo) acelerar() {
	fmt.Println("Acelerando...")
}

func (v Veiculo) freiar() {
	fmt.Println("Freiando...")
}

func (v Veiculo) re() {
	fmt.Println("Dando ré...")
}
