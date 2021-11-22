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

	alias = wtfa.ParseAlias("...=../..")
	assert.NotNil(t, alias)
	assert.Equal(t, "../..", alias.Exec)
	assert.Equal(t, "...", alias.Shortcut)

	alias = wtfa.ParseAlias(`gcm='git checkout $(git_main_branch)'`)
	assert.NotNil(t, alias)
	assert.Equal(t, "git", alias.Exec)
	assert.Equal(t, 2, len(alias.Args))
}
