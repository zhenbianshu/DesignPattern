package main

import (
	"fmt"
	"mediator"
)

func main() {
	france := mediator.NewCountry("France")
	korea := mediator.NewCountry("korea")

	unitedNations := mediator.NewUnitedNations(france, korea)

	france.Announce("I am a romantic country~")
	fmt.Println(korea.Receive(unitedNations)) // I am a romantic country~
	korea.Announce("I am good at copy!")      // I am good at copy!
	fmt.Println(france.Receive(unitedNations))
}
