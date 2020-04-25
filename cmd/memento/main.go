package main

import (
	"fmt"
	me "github.com/yfedoruck/gopattern/pkg/memento"
)

func main() {

	g := me.Git{
		Hash: "qwerty",
	}
	fmt.Println(g)

	m := me.NewMemento(g)

	g.Hash = "asd123"
	fmt.Println(g)
	m.SetState(g)

	g = m.GetState()
	fmt.Println(g)
}
