package main

import (
	"fmt"
	br "github.com/yfedoruck/gopattern/pkg/bridge"
)

func main() {
	win1 := &br.Windows{Archiver: &br.Rar{}}
	win2 := &br.Windows{Archiver: &br.Cab{}}
	lin := &br.Linux{Archiver: &br.Zip{}}

	for _, os := range []br.Os{win1, lin, win2} {
		res := os.Archive("hello, world!")
		fmt.Println(res)
	}
}
