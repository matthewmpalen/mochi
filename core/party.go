package core

import (
	"fmt"
)

type (
	Party struct {
		members []Entity
	}
)

func NewParty(entities ...Entity) Party {
	return Party{members: entities}
}

func (p Party) Members() []Entity {
	return p.members
}

func (p *Party) Damage(x uint64) {
	damagePerMember := x / uint64(len(p.members))
	for i := 0; i < len(p.members); i++ {
		p.members[i].Damage(damagePerMember)
	}
}

func (p Party) IsDefeated() bool {
	for _, entity := range p.members {
		if !entity.IsDead() {
			return false
		}
	}

	return true
}

func (p Party) Size() int {
	return len(p.members)
}

func (p Party) String() string {
	return fmt.Sprintf("Party: IsDefeated=%t", p.IsDefeated())
}
