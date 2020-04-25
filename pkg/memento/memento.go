package memento

import "fmt"

type Git struct {
	Hash string
}

func (r *Git) Commit(hash string) {
	r.Hash = hash
}

type Memento struct {
	state Git
}

func (r *Memento) SetState(g Git) *Memento {
	if r.state.Hash == "" {
		r.state = g
	} else {
		fmt.Println("memento already set")
	}
	return r
}

func (r Memento) GetState() Git {
	return r.state
}

func NewMemento(g Git) Memento {
	m := Memento{}
	m.SetState(g)
	return m
}
