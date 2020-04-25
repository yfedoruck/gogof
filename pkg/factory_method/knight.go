package factory_method

import "fmt"

type Knight struct {
	Man
}
func NewKnight() Warrior {
	k := Knight{}
	k.BuildWeapon()
	return &k
}

func (r *Knight) BuildWeapon() {
	r.Weapon = func() ColdWeapon {
		return Sword{}
	}
}

type Sword struct {
}

func (r Sword) Slash() {
	fmt.Println("Knight hit enemy with Sword!")
}
