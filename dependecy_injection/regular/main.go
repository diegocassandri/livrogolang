package main

import "log"

type Database interface {
	GetUser() string
}

type DefaultDatabse struct {
	name string
}

func (db DefaultDatabse) GetUser() string {
	db.name = "Default Database"
	return db.name
}

type FamousDatabase struct {
	name string
}

func (fb FamousDatabase) GetUser() string {
	fb.name = "Famous Database"
	return fb.name
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s!! Nice to meet you", userName)
}

type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func newProgram(db Database, greeter NiceGreeter) *Program {
	return &Program{Db: db, Greeter: greeter}
}

func main() {
	db := DefaultDatabse{}
	greeter := NiceGreeter{}

	p := newProgram(db, greeter)

	p.Execute()

}
