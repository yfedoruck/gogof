package main

import ch "github.com/yfedoruck/gopattern/pkg/chain"

func main() {
	monkey := ch.Monkey{}
	cat := ch.Cat{}
	dog := ch.Dog{}

	monkey.Next(&cat).Next(&dog)

	for _, food := range []string{"milk", "banana", "grass", "bone"} {
		monkey.Eat(food)
	}

}