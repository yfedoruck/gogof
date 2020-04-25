package chain

import "fmt"

type Dog struct {
	Mammal
}

func (m Dog) Eat(food string) {
	if food == "bone" {
		fmt.Println("Dog gnaws bone")
		return
	}
	if m.next == nil {
		m.theEnd(food)
		return
	}
	(m.next).Eat(food)
}
