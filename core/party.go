package core

type (
	Party struct {
		members []Entity
	}
)

func NewParty(entities ...Entity) Party {
	return Party{
		members: entities,
	}
}

func (p Party) Members() []Entity {
	return p.members
}
