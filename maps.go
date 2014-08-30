package main

import(
  "fmt"
  "strconv"
)

var m map[string]int

func main() {
  go getM()
  //fmt.Println("===========================================")
  getM()
}


func getM(){
  m := make(map[string]int)
  for i:=0;i<10;i++{
    key := strconv.Itoa(i)
    m[key+"_key"] = i
  }

  for i:=0;i<10;i++{
    fmt.Println(m)
  }
}
