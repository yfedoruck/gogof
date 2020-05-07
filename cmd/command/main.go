package main

import (
	"fmt"
	cmd "github.com/yfedoruck/gogof/pkg/command"
	
)

func main() {
	price := 123
	for _, order := range []cmd.Order{cmd.Buy{}, cmd.Sell{}, cmd.Sell{}, cmd.Buy{}} {
		fmt.Println(
			order.Exec(price),
		)
	}
}
