package main

import "fmt"

func main() {
	var (
		fname string
		lname string
	)
	fmt.Println("Olá, Qual o seu nome?")

	n, err := fmt.Scan(&fname, &lname)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("Parameters %d\n", n)
	fmt.Println("hello, " + fname + " " + lname)
}
