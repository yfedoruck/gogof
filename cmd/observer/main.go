package main

import ob "github.com/yfedoruck/gopattern/pkg/observer"

func main() {

	playboy := ob.NewMagazine()
	john := ob.NewHuman("John")
	bob := ob.NewHuman("Bob")
	mary := ob.NewHuman("Mary")

	playboy.Subscribe(john)
	playboy.Subscribe(bob)
	playboy.Subscribe(mary)
	playboy.Publish("New Madonna interview")

	playboy.Unsubscribe(mary)
	playboy.Publish("Angelina Jolie photo")
}
