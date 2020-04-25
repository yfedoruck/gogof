package composite

import "fmt"

type Node interface {
	Data() string
	Display()
}

type Composite struct {
	Title    string
	Children map[string]Node
}

func NewComposite(title string) Composite {
	return Composite{
		Title:    title,
		Children: make(map[string]Node),
	}
}

func (r Composite) Data() string {
	return r.Title
}

func (r *Composite) Add(n Node) {
	r.Children[n.Data()] = n
}

func (r Composite) Remove(title string) {
	if _, ok := r.Children[title]; ok {
		delete(r.Children, title)
	}
}

func (r Composite) Display() {
	fmt.Println(r.Title)
	for _, node := range r.Children {
		node.Display()
	}
}

type Leaf struct {
	Title string
}

func (r Leaf) Data() string {
	return r.Title
}

func (r Leaf) Display() {
	fmt.Println("~" + r.Title + "~")
}
