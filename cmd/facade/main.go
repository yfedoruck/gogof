package main

import fc "github.com/yfedoruck/gogof/pkg/facade"

func main() {
	pc := fc.Computer{
		Su: fc.SystemUnit{},
		Bios: fc.Bios{},
		Os: fc.Windows{},
	}

	pc.On()
	pc.Off()
}
