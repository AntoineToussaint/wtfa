package wtfa

import (
	"regexp"
	"sort"
	"strings"
)

type Alias struct {
	Exec       string
	Definition string
	Shortcut   string
	Args       []string
	SortedArgs []string
}

var aliasRegex *regexp.Regexp

func init() {
	aliasRegex = regexp.MustCompile(`(\w*)='+(\w*)\s+(.*)'+`)
}

type Aliases map[string][]*Alias

func ParseAliases(in []string) (Aliases, int) {
	aliases := make(Aliases)
	count := 0
	for _, line := range in {
		alias := ParseAlias(line)
		if alias == nil {
			continue
		}
		aliases[alias.Exec] = append(aliases[alias.Exec], alias)
		count += 1
	}
	return aliases, count
}

func ParseAlias(s string) *Alias {
	matches := aliasRegex.FindStringSubmatch(s)
	if len(matches) < 2 {
		return nil
	}
	exec := matches[2]
	shortcut := matches[1]
	var args []string
	var sortedArgs []string
	if len(matches) > 3 {
		args = strings.Split(matches[3], " ")
		sortedArgs = strings.Split(matches[3], " ")
	}
	alias := &Alias{
		Exec:       exec,
		Definition: s,
		Shortcut:   shortcut,
		Args:       args,
		SortedArgs: sortedArgs,
	}
	sort.Strings(alias.SortedArgs)
	return alias
}
