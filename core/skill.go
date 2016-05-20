package core

type (
	Skill struct {
		name        string
		description string
	}
)

func NewSkill(n string, desc string) Skill {
	return Skill{name: n, description: desc}
}

func (s Skill) Name() string {
	return s.name
}

func (s Skill) Description() string {
	return s.description
}
