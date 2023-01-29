package guess_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessNumber(t *testing.T) {
	{
		Target = 6
		expected := 5
		n := 10
		assert.NotEqual(t, GuessNumber(n), expected)
	}
	{
		Target = 1
		expected := 1
		n := 2
		assert.Equal(t, GuessNumber(n), expected)
	}
	{
		Target = 1
		expected := 1
		n := 1
		assert.Equal(t, GuessNumber(n), expected)
	}
}
