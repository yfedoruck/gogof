package builder

type BlankPizza struct {
	Pizza Pizza
}

func (r *BlankPizza) NewPizza() {
	r.Pizza = Pizza{}
}

func (r BlankPizza) GetPizza() Pizza {
	return r.Pizza
}

