package main

import (
	"fmt"
	api "github.com/yfedoruck/gogof/pkg/adapter"
)

func main() {
	fb := api.FacebookAdapter{
		Messenger: api.Facebook{},
	}
	tw := api.TwitterAdapter{
		Messenger: api.Twitter{},
	}

	for _, messenger := range []api.Messenger{fb, tw} {
		fmt.Println(
			messenger.Send("Hello World!"),
		)
	}
}
