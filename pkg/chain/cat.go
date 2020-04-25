package chain

import "fmt"

type Cat struct {
	Mammal
}

func (m Cat) Eat(food string) {
	if food == "milk" {
		fmt.Println("Cat drinks milk")
		return
	}
	if m.next == nil {
		m.theEnd(food)
		return
	}
	(m.next).Eat(food)
}
