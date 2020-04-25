package mediator

import (
	"fmt"
	"log"
)

type Db struct {
	Mediator *Mediator
}

func (r Db) Save(data string) {
	fmt.Println("....database try save data....")
	r.Mediator.Listen("db_save", data)
}

type Logger struct {
	Mediator *Mediator
}

func (r Logger) Log(data string) {
	log.Println(data)
}

type Server struct {
	Mediator *Mediator
}

func (r Server) Render(data string) {
	fmt.Println("....server rendering template:.... " + data)
}

func (r Server) Email(data string) {
	fmt.Println("....server sending email....: " + data)
	r.Mediator.Listen("sent_email", data)
}

func (r Server) Controller(data string) {
	fmt.Println("....controller receiving post form....")
	r.Mediator.Listen("post_form", data)
}

type Mediator struct {
	Db     *Db
	Logger *Logger
	Server *Server
}

func (r Mediator) Listen(event string, data string) {
	fmt.Println("event: " + event)
	switch event {
	case "db_save":
		r.Logger.Log("save data to db: " + data)
		r.Server.Render("successfully_saved.html")
	case "post_form":
		r.Logger.Log("receive data: " + data)
	case "sent_email":
		r.Logger.Log("sent email to user")
	}
}

func NewMediator(db *Db, server *Server, logger *Logger) Mediator {
	return Mediator{
		Db:     db,
		Server: server,
		Logger: logger,
	}
}
