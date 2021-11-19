package wtfa

import (
	"regexp"
	"sort"
	"strings"
)

type Cmd struct {
	Full string
	Exec string
	Args []string
}

type Alias struct {
	Definition string
	Shortcut   string
	Args       []string
}

func NewAlias(s string, args []string, def string) *Alias {
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
	return &alias
}

var aliasRegex *regexp.Regexp

func init() {
	aliasRegex = regexp.MustCompile(`(\w*)='(\w*) (.*)'`)
}

type Aliases map[string][]*Alias

func ParseAliases(out string) (Aliases, int) {
	aliases := make(Aliases)
	count := 0
	lines := SplitLines(out)
	for _, line := range lines {
		exec, alias := ParseAlias(line)
		if alias == nil {
			continue
		}
		aliases[exec] = append(aliases[exec], alias)
		count += 1
	}
	return aliases, count
}

func SplitLines(s string) []string {
	var results []string
	current := ""
	inCommand := false
	peek := 0
	for _, c := range s {
		peek += 1
		if c == '\'' && inCommand && (peek < len(s) && s[peek] == ' ') {
			current += string(c)
			inCommand = false
			continue
		}
		if c == '\'' && !inCommand {
			current += string(c)
			inCommand = true
			continue
		}
		if c == ' ' && !inCommand {
			results = append(results, current)
			current = ""
			continue

		}
		current += string(c)
	}
	if current != "" {
		results = append(results, current)
	}
	return results
}

func ParseAlias(s string) (string, *Alias) {
	matches := aliasRegex.FindStringSubmatch(s)
	if len(matches) < 2 {
		return "", nil
	}
	exec := matches[2]
	shortcut := matches[1]
	var args []string
	if len(matches) > 3 {
		args = strings.Split(matches[3], " ")
	}
	return exec, NewAlias(shortcut, args, s)
}
