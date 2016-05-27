package core_test

import (
	"fmt"

	. "github.com/matthewmpalen/mochi/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Party", func() {
	name1 := "test1"
	name2 := "test2"
	var hp uint64 = 10
	var mp uint64 = 9
	var str uint64 = 8
	var def uint64 = 7
	var agil uint64 = 6

	Describe("NewParty", func() {
		It("should return a new Party", func() {
			e1 := NewEntity(name1, hp, mp, str, def, agil)
			e2 := NewEntity(name2, hp, mp, str, def, agil)

			party1 := NewParty(e1)
			Expect(party1.Size()).To(BeEquivalentTo(1))
			Expect(party1.Members()[0]).To(BeEquivalentTo(e1))

			party2 := NewParty(e1, e2)
			Expect(party2.Size()).To(BeEquivalentTo(2))

			members := party2.Members()
			Expect(members[0]).To(BeEquivalentTo(e1))
			Expect(members[1]).To(BeEquivalentTo(e2))
		})
	})

	Describe("Members", func() {
		It("should return the same members", func() {
			expected := []*Entity{
				NewEntity(name1, hp, mp, str, def, agil),
				NewEntity(name2, hp, mp, str, def, agil),
			}

			party1 := NewParty(expected...)
			Expect(party1.Members()).To(BeEquivalentTo(expected))
		})
	})

	Describe("Damage", func() {
		It("should damage all members equally", func() {
			e1 := NewEntity(name1, hp, mp, str, def, agil)
			e2 := NewEntity(name2, hp, mp, str, def, agil)

			party := NewParty(e1, e2)
			party.Damage(2)
			members := party.Members()

			Expect(party.Size()).To(BeEquivalentTo(2))
			Expect(members[0].Stats().HP()).To(BeEquivalentTo(hp - 1))
			Expect(members[1].Stats().HP()).To(BeEquivalentTo(hp - 1))
		})
	})

	Describe("IsDefeated", func() {
		It("should be false if at least one member is not dead", func() {
			e1 := NewEntity(name1, 0, mp, str, def, agil)
			e2 := NewEntity(name2, hp, mp, str, def, agil)

			party := NewParty(e1, e2)
			members := party.Members()

			Expect(members[0].IsDead()).To(BeTrue())
			Expect(members[1].IsDead()).To(BeFalse())
			Expect(party.IsDefeated()).To(BeFalse())
		})

		It("should be true if all members are dead", func() {
			e1 := NewEntity(name1, 0, mp, str, def, agil)
			e2 := NewEntity(name2, 0, mp, str, def, agil)

			party := NewParty(e1, e2)
			members := party.Members()

			Expect(party.Size()).To(BeEquivalentTo(2))
			Expect(members[0].IsDead()).To(BeTrue())
			Expect(members[1].IsDead()).To(BeTrue())
			Expect(party.IsDefeated()).To(BeTrue())
		})
	})

	Describe("Size", func() {
		It("should return 0", func() {
			party := Party{}
			Expect(party.Size()).To(BeEquivalentTo(0))
		})
	})

	Describe("String()", func() {
		It("should return a debug string", func() {
			e1 := NewEntity(name1, 0, mp, str, def, agil)
			e2 := NewEntity(name2, 0, mp, str, def, agil)

			party := NewParty(e1, e2)

			expected := fmt.Sprintf("Party: IsDefeated=%t", party.IsDefeated())
			Expect(party.String()).To(BeEquivalentTo(expected))
		})
	})
})
