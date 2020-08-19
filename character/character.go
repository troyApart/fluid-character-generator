package character

import (
	"fmt"

	"github.com/troyApart/fluid-character-generator/facets/aptitude"
	"github.com/troyApart/fluid-character-generator/facets/skill"
)

var Characters map[int]*Character

const (
	basePCExperience  = 60
	baseNPCExperience = 0
)

type Character struct {
	FirstName  string             `json:"first_name"`
	LastName   string             `json:"last_name"`
	Aptitudes  aptitude.Aptitudes `json:"aptitudes"`
	Skills     skill.Skills       `json:"skills"`
	experience int
}

type CharacterResponse struct {
	Name       string             `json:"name"`
	Experience int                `json:"experience"`
	Aptitudes  aptitude.Aptitudes `json:"aptitudes"`
	Skills     skill.Skills       `json:"skills"`
}

func New(playerCharacter bool) *Character {
	exp := baseNPCExperience
	if playerCharacter {
		exp = basePCExperience
	}
	character := &Character{
		FirstName:  "Some",
		LastName:   "Name",
		experience: exp,
	}
	character.Aptitudes = aptitude.New()
	character.Skills = skill.New()
	character.experience += character.Aptitudes.TotalAptitudePoints() + character.Skills.TotalSkillPoints()

	return character
}

func (c *Character) Get() CharacterResponse {
	var cr CharacterResponse
	if c.FirstName != "" {
		cr.Name = c.FirstName
		if c.LastName != "" {
			cr.Name = fmt.Sprintf("%s %s", cr.Name, c.LastName)
		}
	} else {
		if c.LastName != "" {
			cr.Name = c.FirstName
		}
	}

	cr.Experience = c.Experience()
	cr.Aptitudes = c.Aptitudes
	cr.Skills = c.Skills

	return cr
}

func (c *Character) IncreaseAptitude(s string) error {
	apt, err := aptitude.AptitudeValue(s)
	if err != nil {
		return err
	}

	currentLevel := c.Aptitudes[apt]
	if c.Experience() < aptitude.CalculateAptitudePoints(currentLevel+1) {
		return fmt.Errorf("not enough experience")
	}
	c.Aptitudes[apt] = currentLevel + 1

	return nil
}

func (c *Character) Experience() int {
	exp := c.experience
	exp -= c.Aptitudes.TotalAptitudePoints()
	exp -= c.Skills.TotalSkillPoints()

	return exp
}

func (c *Character) AddExperience(e int) {
	c.experience += e
}
