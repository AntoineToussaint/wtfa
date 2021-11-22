package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCmd(t *testing.T) {
	s := `clear`
	cmd := wtfa.ParseCommand(s)
	assert.Equal(t, 0, len(cmd.Args))
	assert.Equal(t, "clear", cmd.Exec)
}

func TestParseCmdWithQuotes(t *testing.T) {
	s := `git commit -a -m "My commit"`
	cmd := wtfa.ParseCommand(s)
	assert.Equal(t, 4, len(cmd.Args))
	assert.Equal(t, "git", cmd.Exec)
	assert.Equal(t, `"My commit"`, cmd.Args[3])
}

func TestParseDotCmd(t *testing.T) {
	s := `../..`
	cmd := wtfa.ParseCommand(s)
	assert.Equal(t, 0, len(cmd.Args))
	assert.Equal(t, "../..", cmd.Exec)
}
