package core_test

import (
	. "github.com/matthewmpalen/mochi/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Magic", func() {
	Describe("NewMagic", func() {
		It("should return a new Magic", func() {
			name := "Fire"
			description := "It's fire"
			var mp uint64 = 3
			magic := NewMagic(name, description, mp)

			Expect(magic.Name()).To(BeEquivalentTo(name))
			Expect(magic.Description()).To(BeEquivalentTo(description))
			Expect(magic.MP()).To(BeEquivalentTo(mp))
		})
	})
})
