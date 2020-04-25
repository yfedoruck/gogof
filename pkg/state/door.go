package state

type DoorState interface {
	Open() DoorState
	Close() DoorState
	Lock() DoorState
	Unlock() DoorState
}

type Door interface {
	Open()
	Close()
	Lock()
	Unlock()
}

type WoodenDoor struct {
	State DoorState
}

func (r *WoodenDoor) Open() {
	r.State = r.State.Open()
}
func (r *WoodenDoor) Close() {
	r.State = r.State.Close()
}

func (r *WoodenDoor) Lock() {
	r.State = r.State.Lock()
}

func (r *WoodenDoor) Unlock() {
	r.State = r.State.Unlock()
}
