package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	assert.False(t, Match("abcdefdefg", ".*def"))
	assert.False(t, Match("abcdefdefg", ".*def*"))
	assert.False(t, Match("abcdefdefg", ".*fdef"))
	assert.False(t, Match("abcdefdefg", ".*fdef"))

	assert.False(t, Match("abcd", ""))

	assert.True(t, Match("abcd", "a*b*c*d*"))

	assert.True(t, Match("abcdefdef111defdef", "ab.*def.*def"))

	assert.True(t, Match("abcdefdef", "ab.*def"))
	assert.True(t, Match("abcdefdef", ".*def"))

	assert.True(t, Match("", ""))
	assert.True(t, Match("", ".*"))
	assert.False(t, Match("", "."))

	assert.True(t, Match("abcd", "a.*"))
	assert.True(t, Match("abcd", ".*"))
	assert.True(t, Match("abcd", "abcd"))
	assert.True(t, Match("abcd", "abcd*"))
	assert.True(t, Match("abcd", "abc*d*"))
	assert.True(t, Match("abcd", "ab*cd*"))
	assert.True(t, Match("abcd", "abc."))
	assert.True(t, Match("abcd", "ab.d"))
	assert.True(t, Match("abcd", ".bcd"))
	assert.True(t, Match("abcd", ".*cd"))
	assert.True(t, Match("abcd", ".*d"))
	assert.True(t, Match("abcd", "ab.*"))
	assert.True(t, Match("abcd", "abc.*"))
	assert.True(t, Match("abcd", "abcd.*"))

	assert.True(t, Match("abcdef", "abcdef"))
	assert.True(t, Match("abcdef", "ab.def"))
	assert.True(t, Match("abcdef", "ab.*def"))

	assert.True(t, Match("abcdefdef", ".*def.*"))
	assert.True(t, Match("abcdefdef", ".*defdef"))
	assert.True(t, Match("abcdefdef", ".*fdef"))

	assert.True(t, Match("abcdefdefg", ".*fdefg"))
	assert.True(t, Match("abcdefdefg", ".*fdef."))
	assert.True(t, Match("abcdefdefg", ".*fdefg*"))
	assert.True(t, Match("abcdefdefg", ".*def.*"))
	assert.True(t, Match("abcdefdefg", ".*def.*g"))
	assert.True(t, Match("abcdefdefg", ".*defdef.*"))

}
