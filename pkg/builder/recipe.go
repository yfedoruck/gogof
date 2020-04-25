package builder

type PizzaRecipe interface {
	NewPizza()
	Border()
	Filling()
	Seasoning()
	GetPizza() Pizza
}