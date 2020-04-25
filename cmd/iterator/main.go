package main

import (
	"fmt"
	it "github.com/yfedoruck/gopattern/pkg/iterator"
)

func main() {

	var store it.Iterable
	store = &it.Store{
		Fruits: []string{"banana", "apple", "kiwi", "lemon"},
	}

	for store.Exists() {
		f := store.Get()
		fmt.Println(f)

		store.Next()
	}
}
