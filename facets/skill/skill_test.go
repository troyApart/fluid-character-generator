package skill

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkills(t *testing.T) {
	s := New()

	skills := Skills{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0,
	}
	assert.Equal(t, skills, s)

	total := s.TotalSkillPoints()
	assert.Equal(t, 0, total)

	s.RandomSkills()
	total = s.TotalSkillPoints()
	assert.Equal(t, 15, total)
}

func TestSkills_Unmarshal(t *testing.T) {
	b := []byte(`{"melee": 1}`)

	actual := make(Skills, 0)
	err := json.Unmarshal(b, &actual)
	assert.NoError(t, err)

	expectedSkills := make(Skills, 0)
	expectedSkills[Melee] = 1
	assert.Equal(t, expectedSkills, actual)
}
