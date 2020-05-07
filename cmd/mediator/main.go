package main

import me "github.com/yfedoruck/gogof/pkg/mediator"

func main() {

	db := &me.Db{}
	lg := &me.Logger{}
	sr := &me.Server{}
	m := me.NewMediator(db, sr, lg)

	db.Mediator = &m
	lg.Mediator = &m
	sr.Mediator = &m

	db.Save("user credit card")

	sr.Email("** Congratulations! You won 1 000 000 tugriks! **")

}
