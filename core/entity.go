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

func (e *Entity) Heal(x uint64) {
	if e.hp+x > e.maxHP {
		e.hp = e.maxHP
	}

	e.hp += x
	fmt.Printf("%s was healed", e.name)
}

func (e *Entity) Damage(x uint64) {
	if x > e.hp {
		e.hp = 0
	}

	e.hp -= x
	fmt.Printf("%s was lost %d HP", e.name, x)
}

func (e Entity) IsDead() bool {
	return e.hp == 0
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity: name=%s, hp=%d, IsDead=%t", e.name, e.stats.HP(), e.IsDead())
}
