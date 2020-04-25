package builder

import "fmt"

type Director struct {
	recipe PizzaRecipe
}

func (d *Director) Order(order string) *Director {
	switch order {
	case "hawaii":
		d.recipe = &HawaiiPizzaRecipe{}
	case "fish":
		d.recipe = &FishPizzaRecipe{}
	case "vegetarian":
		d.recipe = &VegetarianPizzaRecipe{}
	default:
		fmt.Println("restaurant doesn't have this in menu")
	}
	return d
}

func (d Director) Cook() Pizza {
	d.recipe.NewPizza()
	d.recipe.Filling()
	d.recipe.Border()
	d.recipe.Seasoning()

	return d.recipe.GetPizza()
}
