package wtfa

type Match struct {
	*Cmd
	Aliases []*Alias
}

func FindMatches(cmds []*Cmd, all Aliases) ([]*Match, []*Cmd) {
	var found []*Match
	var unknown []*Cmd
	for _, cmd := range cmds {
		aliases := FindMatch(cmd, all)
		if aliases == nil {
			unknown = append(unknown, cmd)
			continue
		}
		found = append(found, &Match{
			Cmd:     cmd,
			Aliases: aliases,
		})
	}
	return found, unknown
}

func FindMatch(cmd *Cmd, aliases Aliases) []*Alias {
	var matches []*Alias
	uniques := make(map[string]bool)
	for _, alias := range aliases[cmd.Exec] {
		// All arguments match
		if Equal(cmd, alias) {
			return []*Alias{alias}
		}
		if FuzzyMatch(cmd, alias) {
			if _, ok := uniques[alias.Shortcut]; !ok {
				matches = append(matches, alias)
				uniques[alias.Shortcut] = true

			}
		}
	}
	return matches
}

func Equal(cmd *Cmd, alias *Alias) bool {
	if len(cmd.SortedArgs) != len(alias.SortedArgs) {
		return false
	}
	// Verb command should be the same
	if len(cmd.Args) > 0 && cmd.Args[0] != alias.Args[0] {
		return false
	}
	for i, v := range cmd.SortedArgs {
		if v != alias.SortedArgs[i] {
			return false
		}
	}
	return true
}

func OneSubstitutionLengthTwo(cmd *Cmd, alias *Alias) bool {
	// Only when there are more than 2 arguments
	if len(cmd.SortedArgs) <= 1 || len(cmd.SortedArgs) != len(alias.SortedArgs) {
		return false
	}
	// Verb command should be the same
	if cmd.Args[0] != alias.Args[0] {
		return false
	}
	// If we are missing more than 1 substitutions, we fail
	missing := 0
	for _, arg := range cmd.Args {
		if _, ok := alias.MappedArgs[arg]; !ok {
			missing += 1
			if missing > 1 {
				return false
			}
		}
	}
	return true
}

func OneDeletionLengthTwo(cmd *Cmd, alias *Alias) bool {
	// Only when there are more than 2 arguments
	if len(cmd.SortedArgs) <= 1 || len(cmd.SortedArgs) != len(alias.SortedArgs)+1 {
		return false
	}
	// Verb command should be the same
	if cmd.Args[0] != alias.Args[0] {
		return false
	}
	// We should match all arguments
	for _, arg := range alias.Args {
		if _, ok := cmd.MappedArgs[arg]; !ok {
			return false
		}
	}
	return true
}

func FuzzyMatch(cmd *Cmd, alias *Alias) bool {
	// At least 2 arguments and at most one substitution
	if OneSubstitutionLengthTwo(cmd, alias) {
		return true
	}
	if OneDeletionLengthTwo(cmd, alias) {
		return true
	}
	return false
}
