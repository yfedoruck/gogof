package main

import (
	"fmt"
	s "github.com/yfedoruck/gopattern/pkg/strategy"
)

func main() {

	env := "windows"

	archiver := s.Strategy(env)
	pack := archiver.Archive("Hello, World!")
	fmt.Println(pack)

	env = "linux"
	archiver = s.Strategy(env)
	pack = archiver.Archive("Hello, World!")
	fmt.Println(pack)

}


