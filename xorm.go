package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	"os"
	"time"
)

var engine *xorm.Engine

type Venda struct {
	Pedido        int64
	Cliente       int64
	Dataemissao   time.Time
	Valorprodutos float64
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mssql", "server="+os.Getenv("HOST")+";user id="+os.Getenv("USER")+";password="+os.Getenv("PASS")+";database="+os.Getenv("DBNAME"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("conect")
	}
	defer engine.Close()
	engine.ShowSQL = true

	vendas := make([]Venda, 0)
	err = engine.Sql("select top 100 * from Vendas").Find(&vendas)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(vendas)
}
