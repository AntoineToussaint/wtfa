package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCmd(t *testing.T) {
	s := `git commit -a -m "My commit"`
	cmd := wtfa.ParseCommand(s)
	assert.Equal(t, 4, len(cmd.Args))
	assert.Equal(t, "git", cmd.Exec)
	assert.Equal(t, `"My commit"`, cmd.Args[3])
}
