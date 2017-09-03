package main

import (
	"math"
	"math/rand"
)

type Aptitude struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (c *Character) RandomAptitudes() {
	if c.TotalAptitudePoints() < 27 {
		for i := 0; i < 2; i++ {
			aptIndex := rand.Intn(6)
			if c.Aptitudes[aptIndex].Score == 1 {
				c.Aptitudes[aptIndex].Score = 2
			} else {
				i--
			}
		}
		for i := 0; i < 1; i++ {
			aptIndex := rand.Intn(6)
			if c.Aptitudes[aptIndex].Score == 1 && c.Aptitudes[aptIndex].Score != 2 {
				c.Aptitudes[aptIndex].Score = 3
			} else {
				i--
			}
		}
	}
}

func (c *Character) DefaultAptitudeList() {
	aptitudes := []string{"Vigor", "Focus", "Agility", "Cunning", "Perception", "Empathy"}
	c.Aptitudes = make([]Aptitude, 0, 6)
	for _, name := range aptitudes {
		c.Aptitudes = append(c.Aptitudes, Aptitude{
			Name:  name,
			Score: 1,
		})
	}
}

func (c *Character) TotalAptitudePoints() int {
	var exp int
	for _, apt := range c.Aptitudes {
		// fmt.Println("hello", exp)
		exp = exp + calculateAptitudeExperience(apt.Score)
	}
	// fmt.Println("exp", exp)
	return exp
}

func calculateAptitudeExperience(n int) (result int) {
	if n > 1 {
		result = aptitudeScoreToExperience(n) + calculateAptitudeExperience(n-1)
		return result
	}
	return 1
}

func aptitudeScoreToExperience(n int) int {
	return int(math.Pow(float64(n), 2))
}
