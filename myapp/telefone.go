package main

import(
	"strconv"
)
type Telefone struct{
	ddi int
	ddd int
	numero string
}

func (t Telefone) toString() string {
	return "+"+strconv.Itoa(t.ddi)+" ("+strconv.Itoa(t.ddd)+") "+t.numero
}