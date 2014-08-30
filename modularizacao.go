package main

import "fmt"

func main(){
  var a,b,c int
  b = 0
  fmt.Printf("Digite um número: ")
  fmt.Scan(&a)
  fmt.Println("==================")
  fmt.Printf("Tabuada de %v\n",a)
  fmt.Println("==================")
  for b<10{
    b=b+1
    c = mutiplicar(a,b)
    imprimeValor(a,b,c)
  }

  b=0
  fmt.Println("==================")
  for b<10{
    b=b+1
    c = mutiplicar(b,a)
    imprimeValor(b,a,c)
  }
}

func mutiplicar(n1 int, n2 int) int {
  c := n1 * n2
  return c
}

func imprimeValor(n1 int, n2 int, n3 int){
  fmt.Printf("%v x %v = %v\n",n1,n2,n3)
}