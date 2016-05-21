package core_test

import (
	. "github.com/matthewmpalen/mochi/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Skill", func() {
	Describe("NewSkill", func() {
		It("should return a new Skill", func() {
			name := "Skill"
			description := "It's a skill"
			skill := NewSkill(name, description)

			Expect(skill.Name()).To(BeEquivalentTo(name))
			Expect(skill.Description()).To(BeEquivalentTo(description))
		})
	})
})
