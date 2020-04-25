package main

import st "github.com/yfedoruck/gopattern/pkg/state"

func main() {
	var door st.Door
	door = &st.WoodenDoor{State: &st.OpenedDoor{}}
	door.Close()
	door.Lock()
	door.Close()
	door.Lock()
	door.Unlock()
	door.Open()
	door.Open()
	door.Lock()
	door.Close()

}
