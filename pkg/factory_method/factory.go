package factory_method

type Warrior interface {
	BuildWeapon()
	UseWeapon()
}
type ColdWeapon interface {
	Slash()
}
