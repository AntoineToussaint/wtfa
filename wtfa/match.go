package wtfa

func FindMatches(cmds []*Cmd, all Aliases) []*Alias {
	uniques := make(map[string]bool)
	var found []*Alias
	for _, cmd := range cmds {
		alias := FindMatch(cmd, all)
		if alias == nil {
			continue
		}
		if _, exists := uniques[alias.Definition]; exists {
			continue
		}
		found = append(found, alias)
		uniques[alias.Definition] = true
	}
	return found
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func FindMatch(cmd *Cmd, aliases Aliases) *Alias {
	for _, alias := range aliases[cmd.Exec] {
		// All argument match
		if equal(alias.SortedArgs, cmd.SortedArgs) {
			return alias
		}
	}
	return nil
}
