package main

import (
	"fmt"
	"github.com/yfedoruck/gogof/pkg/builder"
)

func main() {

	chief := builder.Director{}

	order := "fish"
	fishPizza := chief.
		Order(order).
		Cook()

	order = "hawaii"
	hawaiiPizza := chief.
		Order(order).
		Cook()

	fmt.Println(fishPizza)
	fmt.Println(hawaiiPizza)

}
