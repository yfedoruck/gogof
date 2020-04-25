package factory_method

import "fmt"

type Samurai struct {
	Man
}

func NewSamurai() Warrior {
	s := Samurai{}
	s.BuildWeapon()
	return &s
}

func (r *Samurai) BuildWeapon() {
	r.Weapon = func() ColdWeapon {
		return Katana{}
	}
}

type Katana struct {
}

func (r Katana) Slash() {
	fmt.Println("Samurai slash enemy with Katana!")
}
