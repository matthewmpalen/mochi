package core

import (
	"math/rand"
)

type (
	Stats struct {
		//base stats
		maxHP        uint64
		maxMP        uint64
		baseStrength uint64
		baseDefense  uint64
		baseAgility  uint64
		baseLuck     uint64
		//derived stats
		hp       uint64
		mp       uint64
		strength uint64
		defense  uint64
		agility  uint64
		luck     uint64
	}
)

func NewStats(hp uint64, mp uint64, str uint64, def uint64, agi uint64) *Stats {
	lck := uint64(rand.Intn(10))
	return &Stats{
		maxHP:        hp,
		hp:           hp,
		maxMP:        mp,
		mp:           mp,
		baseStrength: str,
		strength:     str,
		baseDefense:  def,
		defense:      def,
		baseAgility:  agi,
		agility:      agi,
		baseLuck:     lck,
		luck:         lck,
	}
}

func (s Stats) MaxHP() uint64 {
	return s.maxHP
}

func (s Stats) MaxMP() uint64 {
	return s.maxMP
}

func (s Stats) HP() uint64 {
	return s.hp
}

func (s Stats) MP() uint64 {
	return s.mp
}

func (s Stats) Strength() uint64 {
	return s.strength
}

func (s Stats) Defense() uint64 {
	return s.defense
}

func (s Stats) Agility() uint64 {
	return s.agility
}

func (s Stats) Luck() uint64 {
	return s.luck
}

func (s *Stats) BoostStrength() {
	s.strength *= 2
}

func (s *Stats) BoostDefense() {
	s.defense *= 2
}

func (s *Stats) BoostAgility() {
	s.agility *= 2
}

func (s *Stats) ResetStats() {
	s.hp = s.maxHP
	s.mp = s.maxMP
	s.strength = s.baseStrength
	s.defense = s.baseDefense
	s.agility = s.baseAgility
	s.luck = s.baseLuck
}
