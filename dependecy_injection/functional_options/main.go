package main

import "log"

type Database interface {
	GetUser() string
}

type DefaultDatabase struct {
	name string
}

func (db DefaultDatabase) GetUser() string {
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

// functional option pattern for passing the Interface to the `newProgram()` function
type programOptions struct {
	Db      Database
	Greeter Greeter
}

type programOption func(*programOptions)

func WithDatabase(db Database) programOption {
	return func(po *programOptions) {
		po.Db = db
	}
}

func WithGreeter(g Greeter) programOption {
	return func(po *programOptions) {
		po.Greeter = g
	}
}

// Refactored `newProgram()` function with functional options support
func newProgram(opts ...programOption) *Program {
	options := &programOptions{
		Db:      DefaultDatabase{},
		Greeter: NiceGreeter{},
	}

	for _, opt := range opts {
		opt(options)
	}

	return &Program{
		Db:      options.Db,
		Greeter: options.Greeter,
	}
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

// Example usage:
func main() {
	db := FamousDatabase{}
	greeter := NiceGreeter{}

	p := newProgram(
		WithDatabase(db),
		WithGreeter(greeter),
	)

	p.Execute()

}
