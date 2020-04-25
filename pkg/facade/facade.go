package facade

type Computer struct {
	Su   SystemUnit
	Bios Bios
	Os   Windows
}

func (c Computer) On() {
	c.Su.PowerOn()
	c.Bios.Start()
	c.Os.Start()
}

func (c Computer) Off() {
	c.Os.ShutDown()
	c.Bios.Stop()
	c.Su.PowerOff()
}
