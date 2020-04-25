package bridge

import "fmt"

type Os interface {
	Archive(s string) string
}

type Archiver interface {
	Pack(s string) string
}

type Windows struct {
	Archiver Archiver
}
func (r *Windows) Archive(s string) string {
	fmt.Println("run C:/Program Files/Rar/rar.exe")
	return r.Archiver.Pack(s)
}

type Linux struct {
	Archiver Archiver
}
func (r *Linux) Archive(s string) string {
	fmt.Println("run /bin/rar")
	return r.Archiver.Pack(s)
}

type Rar struct {
}
func (r *Rar) Pack(s string) string {
	return s + " -- has packed with Rar"
}

type Zip struct {
}
func (r *Zip) Pack(s string) string {
	return s + " -- has packed with Zip"
}

type Cab struct {
}
func (r *Cab) Pack(s string) string {
	return s + " -- has packed with CAB"
}