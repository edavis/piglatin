package piglatin

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	var translateTests = []struct {
		in, out string
	}{
		{"dog", "ogday"},
		{"cat", "atcay"},
		{"hat", "athay"},
		{"egg", "eggday"},
	}
	for _, test := range translateTests {
		actual := Translate(test.in)
		assert.Equal(t, test.out, actual)
	}
}
