package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/troyApart/fluid-character-generator/character"
)

type GenerateCharacter struct {
}

func NewGenerateCharacter() *GenerateCharacter {
	character.Characters = make(map[int]*character.Character)
	return &GenerateCharacter{}
}

func (c *GenerateCharacter) Get(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if ch, ok := character.Characters[id]; ok {
		charJSON, _ := json.Marshal(ch.Get())
		w.Write(charJSON)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (c *GenerateCharacter) Post(w http.ResponseWriter, r *http.Request) {
	ch := character.New(true)
	a := ch.Aptitudes.RandomAptitudes()
	s := ch.Skills.RandomSkills()
	ch.AddExperience(a + s)

	character.Characters[len(character.Characters)+1] = ch

	charJSON, _ := json.Marshal(ch.Get())

	w.Write(charJSON)
	w.WriteHeader(http.StatusCreated)
}

func (c *GenerateCharacter) Patch(w http.ResponseWriter, r *http.Request) {
	apt := r.URL.Query().Get("aptitude")
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if ch, ok := character.Characters[id]; ok {
		err := ch.IncreaseAptitude(apt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		charJSON, _ := json.Marshal(ch.Get())
		w.Write(charJSON)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
