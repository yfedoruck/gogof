package state

import "fmt"

type LockedDoor struct {
}

func (r *LockedDoor) Open() DoorState {
	fmt.Println(" ~~ can't open: the door is locked ~~ ")
	return r
}
func (r *LockedDoor) Close() DoorState {
	fmt.Println(" ~~ the door is closed and locked  ~~ ")
	return r
}

func (r *LockedDoor) Lock() DoorState {
	fmt.Println(" ~~ the door already locked ~~ ")
	return r
}

func (r *LockedDoor) Unlock() DoorState {
	fmt.Println("!!  unlock the door  !!")
	return &ClosedDoor{}
}
