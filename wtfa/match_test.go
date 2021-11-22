package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSuccessExactMatch(t *testing.T) {
	cmd := wtfa.ParseCommand("git status")
	alias := wtfa.ParseAlias("gst='git status'")
	assert.True(t, wtfa.Equal(cmd, alias))
}

func TestFindSuccessExactMatchWithoutQuote(t *testing.T) {
	cmd := wtfa.ParseCommand("../..")
	alias := wtfa.ParseAlias("...=../..")
	assert.True(t, wtfa.Equal(cmd, alias))
}

func TestFindSuccessOneSubstitution(t *testing.T) {
	cmd := wtfa.ParseCommand("git checkout master")
	alias := wtfa.ParseAlias("gcm='git checkout $(git_main_branch)'")
	assert.True(t, wtfa.FuzzyMatch(cmd, alias))
}

func TestFindSuccessAliasWithOneDeletion(t *testing.T) {
	cmd := wtfa.ParseCommand("git checkout master")
	alias := wtfa.ParseAlias("gcm='git checkout'")
	assert.True(t, wtfa.FuzzyMatch(cmd, alias))
}
