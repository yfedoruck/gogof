package visitor

import "fmt"

type Visitee interface {
	Visit(v Visitor)
}

type Visitor struct {
}

func (r Visitor) VisitFacebook(fb Facebook) {
	fmt.Println("- visit Facebook and do something here -")
	fmt.Println("Founder of facebook is: " + fb.Founder)
}

func (r Visitor) VisitTwitter(tw Twitter) {
	fmt.Println("-visit Twitter and do something here-")
	fmt.Println("Read Twitter info: " + tw.About)
}

type Facebook struct {
	Founder string
	v       Visitor
}

func (r Facebook) Visit(v Visitor) {
	v.VisitFacebook(r)
}

type Twitter struct {
	About string
	v     Visitor
}

func (r Twitter) Visit(v Visitor) {
	v.VisitTwitter(r)
}
