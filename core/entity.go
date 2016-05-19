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
	if e.stats.hp+x > e.stats.maxHP {
		e.stats.hp = e.stats.maxHP
	}

	e.stats.hp += x
	fmt.Printf("%s was healed", e.name)
}

func (e *Entity) Damage(x uint64) {
	if e.stats.hp == 0 {
		return
	}

	if x > e.stats.hp {
		e.stats.hp = 0
	}

	e.stats.hp -= x
	fmt.Printf("%s lost %d HP (HP=%d, IsDead=%t)\n", e.name, x, e.stats.hp, e.IsDead())
}

func (e Entity) IsDead() bool {
	return e.stats.hp == 0
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity: name=%s, hp=%d, IsDead=%t", e.name, e.stats.hp, e.IsDead())
}
