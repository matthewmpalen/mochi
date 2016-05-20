package core

type (
	Magic struct {
		Skill
		mp uint64
	}
)

func NewMagic(n string, desc string, mp uint64) Magic {
	return Magic{
		Skill: NewSkill(n, desc),
		mp:    mp,
	}
}

func (m Magic) MP() uint64 {
	return m.mp
}
