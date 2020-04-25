package facade

import "fmt"

type Bios struct {
}

func (b Bios) Start() {
	fmt.Println("start BIOS")
}

func (b Bios) Stop() {
	fmt.Println("stopped BIOS")
}
