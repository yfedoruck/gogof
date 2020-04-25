package main

import vs "github.com/yfedoruck/gopattern/pkg/visitor"

func main() {

	v := vs.Visitor{}

	fb := vs.Facebook{Founder: "Zukerberg"}
	fb.Visit(v)

	tw := vs.Twitter{About: "Social platform for birds"}
	tw.Visit(v)

}
