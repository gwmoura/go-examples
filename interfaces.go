package main

import (
	"fmt"
)

func main() {
	var v Veiculo

	c := Carro{3, 4}
	m := Moto{2, 0}

	v = c
	fmt.Print("v agora é e c continua sendo: ")
	fmt.Println(v.nomeVeiculo())

	v = m
	fmt.Print("v agora é e m continua sendo: ")
	fmt.Println(v.nomeVeiculo())

}

// Interface Veiculo
type Veiculo interface {
	setRodas(rodas int)
	setPortas(portas int)
	acelerar()
	freiar()
	trocarMarcha()
	nomeVeiculo() string
}

// Carro
type Carro struct {
	portas int
	rodas  int
}

func (c Carro) setRodas(rodas int) {
	if rodas > 0 && rodas < 5 {
		c.rodas = rodas
	}
}

func (c Carro) setPortas(portas int) {
	if portas > 0 && portas < 5 {
		c.portas = portas
	}
}

func (c Carro) acelerar() {

}

func (c Carro) freiar() {

}

func (c Carro) trocarMarcha() {

}

func (c Carro) nomeVeiculo() string {
	return "carro"
}

// Moto
type Moto struct {
	portas int
	rodas  int
}

func (m Moto) setPortas(portas int) {
	portas = 0
	m.portas = portas
}

func (m Moto) setRodas(rodas int) {
	if rodas > 0 && rodas < 3 {
		m.rodas = rodas
	}
}

func (m Moto) acelerar() {

}

func (m Moto) freiar() {

}

func (m Moto) trocarMarcha() {

}

func (m Moto) nomeVeiculo() string {
	return "moto"
}
