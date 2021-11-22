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
	MappedArgs map[string]bool
}

var aliasOneRegex *regexp.Regexp
var aliasRegex *regexp.Regexp

func init() {
	aliasOneRegex = regexp.MustCompile(`(.*)=(.*)`)
	aliasRegex = regexp.MustCompile(`(\w*)='(\w*)\s?(.*)'`)
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

func parseOneWord(s string) *Alias {
	matches := aliasOneRegex.FindStringSubmatch(s)
	if len(matches) < 3 {
		return nil
	}
	return &Alias{Exec: matches[2], Shortcut: matches[1], Definition: s}
}

func ParseAlias(s string) *Alias {
	matches := aliasRegex.FindStringSubmatch(s)
	if len(matches) < 3 {
		return parseOneWord(s)
	}
	shortcut := matches[1]
	exec := matches[2]
	var args []string
	var sortedArgs []string
	if len(matches) > 3 {
		args = strings.Split(matches[3], " ")
		sortedArgs = strings.Split(matches[3], " ")
	}
	mappedArgs := make(map[string]bool)
	for _, arg := range args {
		mappedArgs[arg] = true
	}
	alias := &Alias{
		Exec:       exec,
		Definition: s,
		Shortcut:   shortcut,
		Args:       args,
		SortedArgs: sortedArgs,
		MappedArgs: mappedArgs,
	}
	sort.Strings(alias.SortedArgs)
	return alias
}
