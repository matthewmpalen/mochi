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

func (e Entity) Name() string {
	return e.name
}

func (e Entity) Stats() Stats {
	return e.stats
}

func (e *Entity) Heal(x uint64) {
	switch {
	case e.IsDead():
		return
	case e.stats.hp == e.stats.maxHP:
		return
	case e.stats.hp+x < e.stats.hp:
		// Overflow
		e.stats.hp = e.stats.maxHP
	case e.stats.hp+x > e.stats.maxHP:
		e.stats.hp = e.stats.maxHP
	default:
		e.stats.hp += x
	}
	fmt.Printf("%s was healed %d", e.name, x)
}

func (e *Entity) Damage(x uint64) {
	switch {
	case e.stats.hp == 0:
		return
	case x > e.stats.hp:
		e.stats.hp = 0
	default:
		e.stats.hp -= x
	}

	fmt.Printf("%s took %d damage (HP=%d, IsDead=%t)\n", e.name, x, e.stats.hp, e.IsDead())
}

func (e Entity) IsDead() bool {
	return e.stats.hp == 0
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity: name=%s, hp=%d, IsDead=%t", e.name, e.stats.hp, e.IsDead())
}
