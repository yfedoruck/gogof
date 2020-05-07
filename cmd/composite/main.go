package main

import (
	cp "github.com/yfedoruck/gogof/pkg/composite"
)

func main() {

	tree := cp.NewComposite("root")
	tree.Add(cp.Leaf{Title: "first"})
	tree.Add(cp.Leaf{Title: "second"})

	branch := cp.NewComposite("branch1")
	branch.Add(cp.Leaf{Title: "branch1-leaf1"})
	branch.Add(cp.Leaf{Title: "branch1-leaf2"})

	tree.Add(branch)

	tree.Display()
}
