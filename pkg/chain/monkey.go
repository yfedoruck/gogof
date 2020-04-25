package chain

import "fmt"

type Monkey struct {
	Mammal
}

func (m Monkey) Eat(food string) {
	if food == "banana" {
		fmt.Println("Monkey eats banana")
		return
	}
	if m.next == nil {
		m.theEnd(food)
		return
	}
	m.next.Eat(food)
}
