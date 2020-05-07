package main

import fm "github.com/yfedoruck/gogof/pkg/factory_method"

func main() {

	warriors := []fm.Warrior{fm.NewKnight(), fm.NewSamurai(), fm.NewAssassin()}
	for _, warrior := range warriors {
		warrior.UseWeapon()
	}

}
