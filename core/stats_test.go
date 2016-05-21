package core_test

import (
	. "github.com/matthewmpalen/mochi/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stats", func() {
	var hp uint64 = 10
	var mp uint64 = 9
	var str uint64 = 8
	var def uint64 = 7
	var agil uint64 = 6

	Describe("NewStats", func() {
		It("should return a new Stats", func() {
			stats := NewStats(hp, mp, str, def, agil)

			Expect(stats.MaxHP()).To(BeEquivalentTo(hp))
			Expect(stats.MaxMP()).To(BeEquivalentTo(mp))
			Expect(stats.HP()).To(BeEquivalentTo(hp))
			Expect(stats.MP()).To(BeEquivalentTo(mp))
			Expect(stats.Strength()).To(BeEquivalentTo(str))
			Expect(stats.Defense()).To(BeEquivalentTo(def))
			Expect(stats.Agility()).To(BeEquivalentTo(agil))
		})
	})

	Describe("Luck", func() {
		It("should return the current luck on a Stats object", func() {
			stats := Stats{}
			Expect(stats.Luck()).To(BeEquivalentTo(0))
		})
	})

	Describe("ResetStats", func() {
		It("should reset the current Stat values to their base values", func() {
			stats := NewStats(hp, mp, str, def, agil)

			stats.BoostStrength()
			stats.BoostDefense()
			stats.BoostAgility()
			stats.ResetStats()

			Expect(stats.Strength()).To(BeEquivalentTo(str))
			Expect(stats.Defense()).To(BeEquivalentTo(def))
			Expect(stats.Agility()).To(BeEquivalentTo(agil))
		})
	})
})
