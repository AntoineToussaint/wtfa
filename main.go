package main

import (
	"fmt"
	"github.com/AntoineToussaint/wtfa/wtfa"
	"os"
	"strings"
)

func main() {
	history := wtfa.GetLastCommand()
	// Pass aliases as argument
	aliases := wtfa.ParseAliases(strings.Join(os.Args[1:], " "))
	match := wtfa.FindMatch(history, aliases)
	if match != nil {
		fmt.Printf("You may try this shortcut: %v\n", match.Definition)
	} else {
		fmt.Println("No alias found")
	}
}
