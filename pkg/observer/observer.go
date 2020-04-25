package observer

import "fmt"

type Subscriber interface {
	Name() string
	Read(msg string)
}

type Magazine struct {
	Subscriber map[string]Subscriber
}

func NewMagazine() Magazine {
	return Magazine{
		Subscriber: make(map[string]Subscriber),
	}
}

func (r Magazine) Publish(msg string) {
	for _, subscriber := range r.Subscriber {
		subscriber.Read(msg)
	}
}

func (r *Magazine) Subscribe(h Subscriber) {
	r.Subscriber[h.Name()] = h
}

func (r *Magazine) Unsubscribe(h Subscriber) {
	delete(r.Subscriber, h.Name())
}

type Human struct {
	name string
}

func (r Human) Name() string {
	return r.name
}

func (r Human) Read(msg string) {
	fmt.Println(r.name + " receives : " + msg)
}

func NewHuman(name string) Human {
	return Human{name: name}
}
