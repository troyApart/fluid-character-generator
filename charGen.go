package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const (
	basePCExperience       = 60
	baseCreatureExperience = 0
)

type Character struct {
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Experience int        `json:"experience"`
	Aptitudes  []Aptitude `json:"aptitudes"`
	Skills     []Skill    `json:"skills"`
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	character := GenerateCharacter(true)
	character.DefaultAptitudeList()
	character.DefaultSkillList()
	// fmt.Println(character.TotalAptitudePoints())

	character.RandomAptitudes()
	character.RandomSkills()
	// c := fmt.Sprintf("%#v", character)
	// fmt.Printf("%+v", character)
	charJSON, _ := json.Marshal(character)
	fmt.Println(string(charJSON))
	// fmt.Println("")
	// fmt.Println(character.TotalAptitudePoints())
	// fmt.Println(rand.Intn(6))
}

func GenerateCharacter(playerCharacter bool) *Character {
	exp := baseCreatureExperience
	if playerCharacter {
		exp = basePCExperience
	}
	return &Character{
		FirstName:  "Some",
		LastName:   "Name",
		Experience: exp,
	}
}
