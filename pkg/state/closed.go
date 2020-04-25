package state

import "fmt"

type ClosedDoor struct {
}

func (r *ClosedDoor) Open() DoorState {
	fmt.Println("!!  open the door  !!")
	return &OpenedDoor{}
}
func (r *ClosedDoor) Close() DoorState {
	fmt.Println(" ~~ door already closed ~~ ")
	return r
}

func (r *ClosedDoor) Lock() DoorState {
	fmt.Println("!!  lock the door  !!")
	return &LockedDoor{}
}

func (r *ClosedDoor) Unlock() DoorState {
	fmt.Println(" ~~ the door is not  locked ~~ ")
	return r
}
