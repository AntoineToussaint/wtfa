package wtfa

import (
	"reflect"
)

func FindMatch(cmd Cmd, all Aliases) *Alias {
	if aliases, ok := all[cmd.Exec]; ok {
		for _, alias := range aliases {
			if reflect.DeepEqual(cmd.Args, alias.Args) {
				return alias
			}
		}
		// Match but one
		for _, alias := range aliases {
			if MatchButOne(cmd.Args, alias.Args) {
				return alias
			}
		}

	}
	return nil
}

func MatchButOne(cmd, alias []string) bool {
	// if the alias is too complex, we skip
	if len(alias) > len(cmd) {
		return false
	}
	// if the alias is too short, we skip
	if len(alias) < len(cmd)-1 {
		return false
	}
	catch := make(map[string]bool)
	for _, l := range cmd {
		catch[l] = true
	}
	mismatch := 0
	for _, r := range alias {
		if _, ok := catch[r]; !ok {
			mismatch += 1
			if mismatch > 1 || len(alias) == 1 {
				return false
			}
		}
	}
	return true
}
