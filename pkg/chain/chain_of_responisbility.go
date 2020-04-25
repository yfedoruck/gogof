package chain

import "fmt"

type Animal interface {
	Eat(food string)
	Next(animal Animal) Animal
}

type Mammal struct {
	next Animal
}

func (m *Mammal) Next(animal Animal) Animal {
	m.next = animal
	return m.next
}

func (m Mammal) theEnd(food string) {
	fmt.Println("no one eats it: " + food)
}
