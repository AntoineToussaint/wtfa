package wtfa

import (
	"sort"
	"strings"
)

type Cmd struct {
	Full       string
	Exec       string
	Args       []string
	SortedArgs []string
	MappedArgs map[string]bool
}

func ParseCommand(full string) *Cmd {
	tokens := strings.Split(full, " ")
	cmd := Cmd{Full: full, Exec: tokens[0], MappedArgs: make(map[string]bool)}
	args := ParseArgs(strings.Join(tokens[1:], " "))
	for _, arg := range args {
		if arg != "" {
			cmd.Args = append(cmd.Args, arg)
			cmd.SortedArgs = append(cmd.SortedArgs, arg)
			cmd.MappedArgs[arg] = true
		}
	}
	sort.Strings(cmd.SortedArgs)
	return &cmd
}

func ParseArgs(s string) []string {
	inQuote := false
	var args []string
	current := ""
	for _, c := range s {
		if c == ' ' && !inQuote {
			args = append(args, current)
			current = ""
			continue
		}
		if c == ' ' && inQuote {
			current = current + string(c)
			continue
		}
		if c == '"' && inQuote {
			current = current + string(c)
			inQuote = false
			continue
		}
		if c == '"' && !inQuote {
			current = current + string(c)
			inQuote = true
			continue
		}
		current = current + string(c)
	}
	if len(current) > 0 {
		args = append(args, current)
	}
	return args
}
