package aptitude

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAptitude(t *testing.T) {
	a := New()

	aptitudes := Aptitudes{
		0: 1, 1: 1, 2: 1, 3: 1, 4: 1, 5: 1,
	}
	assert.Equal(t, aptitudes, a)

	total := a.TotalAptitudePoints()
	assert.Equal(t, 6, total)

	a.RandomAptitudes()
	total = a.TotalAptitudePoints()
	assert.Equal(t, 27, total)
}

func TestAptitudes_Unmarshal(t *testing.T) {
	b := []byte(`{"agility": 1}`)

	actual := make(Aptitudes, 0)
	err := json.Unmarshal(b, &actual)
	assert.NoError(t, err)

	expectedSkills := make(Aptitudes, 0)
	expectedSkills[Agility] = 1
	assert.Equal(t, expectedSkills, actual)
}
