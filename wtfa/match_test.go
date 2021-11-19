package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSuccessMatch(t *testing.T) {
	history := wtfa.NewCmd("git", []string{"add", "."})
	aliases := make(wtfa.Aliases)
	aliases["git"] = append(aliases["git"], wtfa.NewAlias("gca", []string{"add"}))
	assert.NotNil(t, wtfa.FindMatch(history, aliases))
}

func TestFindFailMatch(t *testing.T) {
	history := wtfa.NewCmd("git", []string{"add", "."})
	aliases := make(wtfa.Aliases)
	aliases["git"] = append(aliases["git"], wtfa.NewAlias("g", nil))
	assert.Nil(t, wtfa.FindMatch(history, aliases))
}
