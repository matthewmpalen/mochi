package core

import "fmt"

type (
	Entity struct {
		name  string
		stats Stats
	}
)

func NewEntity(n string, hp uint64, mp uint64, str uint64, def uint64, a uint64) Entity {
	return Entity{
		name:  n,
		stats: NewStats(hp, mp, str, def, a),
	}
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity: name=%s, hp=%d", e.name, e.stats.HP())
}
