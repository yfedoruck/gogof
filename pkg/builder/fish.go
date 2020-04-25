package builder

type FishPizzaRecipe struct {
	BlankPizza
}

func (r *FishPizzaRecipe) Border() {
	r.Pizza.Border = "fish"
}
func (r *FishPizzaRecipe) Filling() {
	r.Pizza.Filling = "crab sticks"
}
func (r *FishPizzaRecipe) Seasoning() {
	r.Pizza.Seasoning = "sea salt"
}
