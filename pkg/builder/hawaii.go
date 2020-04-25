package builder

type HawaiiPizzaRecipe struct {
	BlankPizza
}

func (r *HawaiiPizzaRecipe) Border() {
	r.Pizza.Border = "cheese"
}
func (r *HawaiiPizzaRecipe) Filling() {
	r.Pizza.Filling = "chicken and pineapple"
}
func (r *HawaiiPizzaRecipe) Seasoning() {
	r.Pizza.Seasoning = "sugar"
}
