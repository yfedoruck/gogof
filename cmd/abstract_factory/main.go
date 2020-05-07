package main

import (
	"fmt"
	af "github.com/yfedoruck/gogof/pkg/abstract_factory"
)

func main() {
	win := af.Windows{}
	lin := af.Linux{}

	for _, os := range []af.Os{win, lin} {
		archiver := os.GetArchiver()
		hw := archiver.Archive("Hello, World!")
		fmt.Println(hw)
	}
}
