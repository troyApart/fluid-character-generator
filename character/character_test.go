package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	c := New(true)
	assert.Equal(t, "Some", c.FirstName)
	assert.Equal(t, "Name", c.LastName)
}
