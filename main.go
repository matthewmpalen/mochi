package main

import (
	"fmt"

	"github.com/matthewmpalen/mochi/core"
)

func main() {
	e1 := core.NewEntity("A", 100, 50, 5, 5, 5)
	e2 := core.NewEntity("B", 80, 70, 4, 4, 6)
	party := core.NewParty(e1, e2)

	for !party.IsDefeated() {
		party.Damage(10)
		fmt.Println(party)
	}
}
