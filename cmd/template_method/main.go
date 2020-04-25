package main

import (
	"fmt"
	tm "github.com/yfedoruck/gopattern/pkg/template_method"
)

func main() {

	for _, archiver := range []tm.Archiver{tm.NewRar(), tm.NewZip()} {
		packed := archiver.Pack("Hello, World!")
		fmt.Println(packed)
	}
}
