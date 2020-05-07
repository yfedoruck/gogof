package main

import (
	txt "github.com/yfedoruck/gogof/pkg/decorator"
	"log"
)

func main() {
	p := txt.Plain{}
	b := txt.Bold{
		Text: p,
	}
	i := txt.Italic{
		Text: p,
	}
	bi := txt.Bold{
		Text: i,
	}

	for _, text := range []txt.Text{p, b, i, bi} {
		log.Println(text.Format("Hello World!"))
	}

}
