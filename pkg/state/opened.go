package state

import "fmt"

type OpenedDoor struct {
}

func (r *OpenedDoor) Open() DoorState {
	fmt.Println(" ~~ the door already opened ~~ ")
	return r
}
func (r *OpenedDoor) Close() DoorState {
	fmt.Println("!!  close the door  !!")
	return &ClosedDoor{}
}

func (r *OpenedDoor) Lock() DoorState {
	fmt.Println(" ~~ can't lock: the door is opened ~~ ")
	return r
}

func (r *OpenedDoor) Unlock() DoorState {
	fmt.Println(" ~~ the door already unlocked ~~ ")
	return r
}
