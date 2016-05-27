package battle_test

import (
	"fmt"

	. "github.com/matthewmpalen/mochi/battle"
	"github.com/matthewmpalen/mochi/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Turn", func() {
	entity := core.NewEntity("test", 10, 10, 10, 10, 10)
	party := core.NewParty(entity)

	Describe("NewTurn", func() {
		It("should return a new Turn", func() {
			var count uint = 1
			turn := NewTurn(count, party)
			expectedString := fmt.Sprint("Turn %d", turn.Count())

			Expect(turn.Count()).To(BeEquivalentTo(count))
			Expect(turn.Party()).To(BeEquivalentTo(party))
			Expect(turn.String()).To(BeEquivalentTo(expectedString))
		})
	})
})
