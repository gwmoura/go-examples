package main

import "fmt"

func main(){
  var a,b,c int

  fmt.Printf("Digite um número: ")
  fmt.Scan(&a)
  fmt.Println("==================")
  fmt.Printf("Tabuada de %v\n",a)
  fmt.Println("==================")
  for b<10{
    b=b+1
    c = a*b
    fmt.Printf("%v x %v = %v\n",a,b,c)
  }
}