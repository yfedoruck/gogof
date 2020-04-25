package factory_method

type Man struct {
	Weapon func() ColdWeapon
}

func (r Man) UseWeapon() {
	weapon := r.Weapon()
	weapon.Slash()
}