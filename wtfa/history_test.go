package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArgs(t *testing.T) {
	s := `commit -m "My message"`
	assert.Equal(t, []string{"commit", "-m", `"My message"`}, wtfa.ParseArgs(s))
}

func TestParse(t *testing.T) {
	s := ": 1637273587:0;git add ."
	expected := wtfa.Cmd{Exec: "git", Args: []string{".", "add"}}
	got := wtfa.ParseZshHistory(s)
	assert.Equal(t, expected.Exec, got.Exec)
	assert.Equal(t, expected.Args, got.Args)
}
