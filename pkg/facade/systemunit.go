package facade

import "fmt"

type SystemUnit struct {
}

func (c SystemUnit) PowerOn() {
	fmt.Println("Button `start` pressed")
}

func (c SystemUnit) PowerOff() {
	fmt.Println("The power is off")
}