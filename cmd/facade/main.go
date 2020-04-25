package main

import fc "github.com/yfedoruck/gopattern/pkg/facade"

func main() {
	pc := fc.Computer{
		Su: fc.SystemUnit{},
		Bios: fc.Bios{},
		Os: fc.Windows{},
	}

	pc.On()
	pc.Off()
}
