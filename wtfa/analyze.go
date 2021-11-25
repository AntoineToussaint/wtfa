package wtfa

type Proposal struct {
	Count int
	Shortcut string
	Full string
}

func Analyze(unknowns []*Cmd, stats Stats, aliases Aliases) []*Proposal{
	// Extract all shortcuts
	shortcuts := make(map[string]bool)
	for _, list := range aliases {
		for _, alias := range list {
			shortcuts[alias.Shortcut] = true
		}
	}
	var proposals []*Proposal
	for _, cmd  := range unknowns {
		if _, ok := shortcuts[cmd.Full]; ok {
			continue
		}
		if len(cmd.Full) < 3 {
			// Probably not worth it
			continue
		}
		if stats[cmd.Full] > 10 {
			proposal := CreateProposal(cmd, shortcuts)
			if proposal == nil {
				continue
			}
			proposal.Count = stats[cmd.Full]
			proposals = append(proposals, proposal)
		}
	}
	return proposals
}

func CreateProposal(cmd *Cmd, shortcuts map[string]bool) *Proposal {
	return &Proposal{
		Shortcut: "TBD",
		Full:     cmd.Full,
	}
}


