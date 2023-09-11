package airportrobot

import (
	"fmt"
)

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeting Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeting.LanguageName(), 	greeting.Greet(name))
}

type Italian struct {}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

type Portuguese struct {}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}