package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		input   string
		ignores []rune
		expect  string
	}{
		{``, nil, ``},
		{`abcd`, nil, `abcd`},
		{`ab,cd,ef`, []rune{','}, `abcdef`},
		{`ab, cd,ef`, []rune{',', ' '}, `abcdef`},
		{`ab, cd, ef`, []rune{',', ' '}, `abcdef`},
	}

	for _, each := range testCases {
		t.Run(each.input, func(t *testing.T) {
			actual := Filter(each.input, func(r rune) bool {
				for _, x := range each.ignores {
					if x == r {
						return true
					}
				}
				return false
			})
			assert.Equal(t, each.expect, actual)
		})
	}
}
