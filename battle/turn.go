package battle

import (
	"fmt"

	"github.com/matthewmpalen/mochi/core"
)

type (
	Turn struct {
		count uint
		party *core.Party
	}
)

func NewTurn(c uint, p *core.Party) *Turn {
	return &Turn{count: c, party: p}
}

func (t Turn) Count() uint {
	return t.count
}

func (t Turn) Party() *core.Party {
	return t.party
}

func (t Turn) String() string {
	return fmt.Sprint("Turn %d", t.count)
}
