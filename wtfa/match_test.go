package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSuccessMatch(t *testing.T) {
	history := wtfa.NewCmd("git", []string{"add", "."})
	aliases := make(wtfa.Aliases)
	aliases["git"] = append(aliases["git"], wtfa.NewAlias("gca", []string{"add"}, "def"))
	assert.NotNil(t, wtfa.FindMatch(history, aliases))
}

func TestFindFailMatchTooShort(t *testing.T) {
	history := wtfa.NewCmd("git", []string{"add", "."})
	aliases := make(wtfa.Aliases)
	aliases["git"] = append(aliases["git"], wtfa.NewAlias("g", nil, "def"))
	assert.Nil(t, wtfa.FindMatch(history, aliases))
}

func TestFindFailMatch(t *testing.T) {
	history := wtfa.NewCmd("git", []string{"push", "origin"})
	aliases := make(wtfa.Aliases)
	aliases["git"] = append(aliases["git"], wtfa.NewAlias("ga", []string{"add"}, "def"))
	assert.Nil(t, wtfa.FindMatch(history, aliases))
}
