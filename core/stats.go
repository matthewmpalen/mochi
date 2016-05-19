package core

type (
	Stats struct {
		hp       uint64
		mp       uint64
		strength uint64
		defense  uint64
		agility  uint64
	}
)

func NewStats(hp uint64, mp uint64, str uint64, def uint64, a uint64) Stats {
	return Stats{
		hp:       hp,
		mp:       mp,
		strength: str,
		defense:  def,
		agility:  a,
	}
}

func (s Stats) HP() uint64 {
	return s.hp
}
