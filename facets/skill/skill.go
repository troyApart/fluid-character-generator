package skill

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

const basePCSkillPoints = 0

type Skill int

const (
	Melee Skill = iota
	Unarmed
	Ranged
	Athletics
	Acrobatics
	Persuasion
	Stealth
)

func SkillValue(s string) (Skill, error) {
	switch strings.ToLower(s) {
	case "melee":
		return Melee, nil
	case "unarmed":
		return Unarmed, nil
	case "ranged":
		return Ranged, nil
	case "athletics":
		return Athletics, nil
	case "acrobatics":
		return Acrobatics, nil
	case "persuasion":
		return Persuasion, nil
	case "stealth":
		return Stealth, nil
	}
	return 0, fmt.Errorf("not found")
}

type Skills map[Skill]int

func (d Skill) String() string {
	return [...]string{"Melee", "Unarmed", "Ranged", "Athletics", "Acrobatics", "Persuasion", "Stealth"}[d]
}

func New() Skills {
	s := make(map[Skill]int, 7)
	for i := Skill(0); i < 7; i++ {
		s[i] = 0
	}
	return s
}

func (s Skills) RandomSkills() int {
	if s.TotalSkillPoints() == basePCSkillPoints {
		for i := 0; i < 3; i++ {
			skillIndex := rand.Intn(7)
			if s[Skill(skillIndex)] == 0 {
				s[Skill(skillIndex)] = 1
			} else {
				i--
			}
		}
		for i := 0; i < 2; i++ {
			skillIndex := rand.Intn(7)
			if s[Skill(skillIndex)] == 0 {
				s[Skill(skillIndex)] = 2
			} else {
				i--
			}
		}
		for i := 0; i < 1; i++ {
			skillIndex := rand.Intn(7)
			if s[Skill(skillIndex)] == 0 {
				s[Skill(skillIndex)] = 3
			} else {
				i--
			}
		}
	} else {
		fmt.Println("not starting from base skills")
	}

	return s.TotalSkillPoints() - basePCSkillPoints
}

func (s Skills) TotalSkillPoints() int {
	var total int
	for _, score := range s {
		for i := score; i > 0; i-- {
			total += i
		}
	}
	return total
}

func (s Skills) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(s)
	count := 0

	keys := make([]int, 0, len(s))
	for key := range s {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		buffer.WriteString(fmt.Sprintf("\"%s\":%d", strings.ToLower(Skill(key).String()), s[Skill(key)]))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (a Skills) UnmarshalJSON(b []byte) error {
	var skills map[string]int
	err := json.Unmarshal(b, &skills)
	if err != nil {
		return err
	}
	for skill, score := range skills {
		skillID, err := SkillValue(skill)
		if err != nil {
			return err
		}
		a[skillID] = score
	}
	return nil
}
