package facade

import "fmt"

type Windows struct {
}

func (w Windows) Start() {
	fmt.Println("Windows started")
}

func (w Windows) ShutDown() {
	fmt.Println("Windows shut down")
}

