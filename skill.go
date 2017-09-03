package main

import "math/rand"

type Skill struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (c *Character) DefaultSkillList() {
	skills := []string{"Melee", "Unarmed", "Ranged", "Athletics", "Acrobatics", "Persuasion", "Stealth"}
	c.Skills = make([]Skill, 0, 6)
	for _, name := range skills {
		c.Skills = append(c.Skills, Skill{
			Name:  name,
			Score: 0,
		})
	}
}

func (c *Character) RandomSkills() {
	if c.TotalSkillPoints() < 15 {
		for i := 0; i < 3; i++ {
			skillIndex := rand.Intn(6)
			if c.Skills[skillIndex].Score == 0 {
				c.Skills[skillIndex].Score = 1
			} else {
				i--
			}
		}
		for i := 0; i < 2; i++ {
			skillIndex := rand.Intn(6)
			if c.Skills[skillIndex].Score == 0 && c.Skills[skillIndex].Score != 1 {
				c.Skills[skillIndex].Score = 2
			} else {
				i--
			}
		}
		for i := 0; i < 1; i++ {
			skillIndex := rand.Intn(6)
			if c.Skills[skillIndex].Score == 0 && c.Skills[skillIndex].Score != 1 && c.Skills[skillIndex].Score != 2 {
				c.Skills[skillIndex].Score = 3
			} else {
				i--
			}
		}
	}
}

func (c *Character) TotalSkillPoints() int {
	var exp int
	for _, skill := range c.Skills {
		exp = exp + calculateSkillExperience(skill.Score)
	}
	return exp
}

func calculateSkillExperience(n int) (result int) {
	if n > 1 {
		result = n + calculateAptitudeExperience(n-1)
		return result
	}
	return 1
}
