package main

import(
	"fmt"
)

func main() {
	var p Pessoa
	p.first_name = "George"
	p.email = "gwmoura@gmail.com"
	p.Telefone=Telefone{55,71,"8788-6089"}
	fmt.Println("Dados da Pessoa:\n")
	fmt.Println(p.toString())
}