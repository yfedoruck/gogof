package main

import (
	"fmt"
	pr "github.com/yfedoruck/gogof/pkg/prototype"
)

func main() {

	cinema := pr.Cinema{
		Title: "Titanic",
		DownloadedFile: "Titanic.mkv",
	}

	prototype := pr.Prototype{Cinema: cinema}

	copy1 := prototype.NewCinemaCopy()
	copy2 := prototype.NewCinemaCopy()


	fmt.Println(&copy1 != &copy2)
}
