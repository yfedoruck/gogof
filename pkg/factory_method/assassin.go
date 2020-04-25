package factory_method

import "fmt"

type Assassin struct {
	Man
}

func NewAssassin() Warrior {
	a := Assassin{}
	a.BuildWeapon()
	return &a
}

func (r *Assassin) BuildWeapon() {
	r.Weapon = func() ColdWeapon {
		return Knife{}
	}
}

type Knife struct {
}

func (r Knife) Slash() {
	fmt.Println("Assassin strike enemy with Knife!")
}
