package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSuccessExactMatch(t *testing.T) {
	cmd := wtfa.ParseCommand("git status")
	assert.Equal(t, "git", cmd.Exec)
	aliases := make(wtfa.Aliases)
	alias := wtfa.ParseAlias("gst=git status")
	assert.NotNil(t, alias)
	aliases["git"] = append(aliases["git"], alias)
	assert.NotNil(t, wtfa.FindMatch(cmd, aliases))
}
