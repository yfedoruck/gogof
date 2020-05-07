package main

import (
	"fmt"
	"github.com/yfedoruck/gogof/pkg/proxy"
)

func main() {

	img1 := proxy.NewImage()
	fmt.Println(img1)

	img1.Show()
	fmt.Println(img1)

	img2 := proxy.NewImage()
	fmt.Println(img2)
	img2.Show()
	fmt.Println(img2)

	fmt.Println(img2 == img1)
}
