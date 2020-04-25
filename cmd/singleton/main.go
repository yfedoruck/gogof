package main

import (
	"fmt"
	"github.com/yfedoruck/gopattern/pkg/singleton"
)

func main() {
	db1 := singleton.Database("connection-data1")
	db2 := singleton.Database("connection-data2")
	fmt.Println(db1.Connect)
	fmt.Println(db2.Connect)
	fmt.Println(db1 == db2)
}
