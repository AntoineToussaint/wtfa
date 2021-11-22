package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAlias(t *testing.T) {
	alias := wtfa.ParseAlias("gst='git status'")
	assert.NotNil(t, alias)
	assert.Equal(t, "git", alias.Exec)
	assert.Equal(t, "gst", alias.Shortcut)

	alias = wtfa.ParseAlias("rb=ruby")
	assert.NotNil(t, alias)
	assert.Equal(t, "ruby", alias.Exec)
	assert.Equal(t, "rb", alias.Shortcut)
}
