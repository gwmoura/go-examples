package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Agora")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Passou um segundo")
	}

}
