package proxy

import "fmt"

type Image interface {
	Show()
}

type ImageHD struct {
	Src string
}

func (r ImageHD) Show()  {
	fmt.Println("Show FullHD Image")
}


type Placeholder struct {
	Src string
}

func (r Placeholder) Show()  {
	fmt.Println("show placeholder")
}


type Proxy struct {
	img Image
}

func (r *Proxy) Show()  {
	r.img = ImageHD{"Download very large file"}
	r.img.Show()
}

func NewImage() Image {
	return &Proxy{Placeholder{"dummy text"}}
}
