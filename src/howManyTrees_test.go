package src

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var n, h = 3, 2
var n2, h2 = 3, 3

func Test1(t *testing.T) {
	excpect := 5
	assert.Equal(t, Solution(n, h), excpect)

}

func Test2(t *testing.T) {
	excpect := 4
	assert.Equal(t, Solution(n2, h2), excpect)
}
