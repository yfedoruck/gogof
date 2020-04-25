package builder

type VegetarianPizzaRecipe struct {
	BlankPizza
}

func (r *VegetarianPizzaRecipe) Border() {
	r.Pizza.Border = "tomato"
}
func (r *VegetarianPizzaRecipe) Filling() {
	r.Pizza.Filling = "apple"
}
func (r *VegetarianPizzaRecipe) Seasoning() {
	r.Pizza.Seasoning = "pepper"
}