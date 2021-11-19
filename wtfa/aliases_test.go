package wtfa_test

import (
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitLines(t *testing.T) {
	s := "7='cd -7' 8='cd -8' 9='cd -9' RED='RAILS_ENV=development' REP='RAILS_ENV=production'"
	lines := wtfa.SplitLines(s)
	assert.Equal(t, 5, len(lines))
	assert.Equal(t, "8='cd -8'", lines[1])

	s = "-='cd -' ...=../.. 1='cd -1'"
	lines = wtfa.SplitLines(s)
	assert.Equal(t, 3, len(lines))
	assert.Equal(t, "1='cd -1'", lines[2])
}
