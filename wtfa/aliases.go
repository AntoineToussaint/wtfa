package wtfa

import (
	"regexp"
	"sort"
	"strings"
)

type Cmd struct {
	Exec string
	Args []string
}

type Alias struct {
	Definition string
	Shortcut   string
	Args       []string
}

func NewAlias(s string, args []string, def string) Alias {
	alias := Alias{
		Shortcut:   s,
		Definition: def,
	}
	for _, arg := range args {
		if arg != "" {
			alias.Args = append(alias.Args, arg)
		}
	}

	sort.Strings(alias.Args)
	return alias
}

var aliasRegex *regexp.Regexp

func init() {
	aliasRegex = regexp.MustCompile(`(.*)=(\w*)\s*(\w*)`)
}

type Aliases map[string][]Alias

func ParseAliases(out string) Aliases {
	aliases := make(Aliases)
	lines := strings.Split(out, "\\n")[1:]
	lines = lines[0 : len(lines)-1]
	for _, line := range lines {
		line = strings.Replace(line, "\\'", "", -1)
		exec, alias := ParseAlias(line)
		aliases[exec] = append(aliases[exec], alias)
	}
	return aliases
}

func ParseAlias(s string) (string, Alias) {
	matches := aliasRegex.FindStringSubmatch(s)
	exec := matches[2]
	shortcut := matches[1]
	var args []string
	if len(matches) > 3 {
		args = strings.Split(matches[3], " ")
	}
	return exec, NewAlias(shortcut, args, s)
}
