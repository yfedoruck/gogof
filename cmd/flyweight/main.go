package main

import (
	"fmt"
	fw "github.com/yfedoruck/gogof/pkg/flyweight"
)

func main() {

	ab := fw.NewAlphabet()
	b := ab.GetLetter("b")
	bCopy := ab.GetLetter("b")

	fmt.Println(b.GetSign(), b == bCopy)
}
