package battle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBattle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Battle Suite")
}
