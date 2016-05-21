package core_test

import (
	"math"

	. "github.com/matthewmpalen/mochi/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Entity", func() {
	name := "test"
	var hp uint64 = 10
	var mp uint64 = 9
	var str uint64 = 8
	var def uint64 = 7
	var agil uint64 = 6

	Describe("NewEntity", func() {
		It("should return a new Entity", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			Expect(entity.Name()).To(BeEquivalentTo(name))
		})
	})

	Describe("Heal", func() {
		It("should not heal a dead Entity", func() {
			entity := NewEntity(name, 0, mp, str, def, agil)
			entity.Heal(1)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(0))
		})

		It("should heal an Entity by 0", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			var previousHP uint64 = entity.Stats().HP()
			entity.Heal(0)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(previousHP))
		})

		It("should heal an Entity by 1", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			entity.Damage(2)
			var previousHP uint64 = entity.Stats().HP()
			entity.Heal(1)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(previousHP + 1))
		})

		It("should heal an Entity by more than maxHP - hp", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			entity.Damage(1)
			entity.Heal(2)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(entity.Stats().MaxHP()))
		})

		It("should heal an Entity by MaxUint64 and cap at maxHP", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			entity.Damage(1)
			entity.Heal(math.MaxUint64)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(entity.Stats().MaxHP()))
		})
	})

	Describe("Damage", func() {
		It("should not damage a dead Entity", func() {
			entity := NewEntity(name, 0, mp, str, def, agil)
			entity.Damage(1)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(0))
		})

		It("should damage an Entity by 0", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			var previousHP uint64 = entity.Stats().HP()
			entity.Damage(0)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(previousHP))
		})

		It("should damage an Entity by 1", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			var previousHP uint64 = entity.Stats().HP()
			entity.Damage(1)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(previousHP - 1))
		})

		It("should damage an Entity by no more than hp", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			var previousHP uint64 = entity.Stats().HP()
			entity.Damage(previousHP + 1)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(0))
		})

		It("should damage an Entity by MaxUint64 and set hp to 0", func() {
			entity := NewEntity(name, hp, mp, str, def, agil)
			entity.Damage(math.MaxUint64)
			Expect(entity.Stats().HP()).To(BeEquivalentTo(0))
		})
	})
})
