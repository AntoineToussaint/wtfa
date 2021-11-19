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
	aliases, count := wtfa.ParseAliases(strings.Join(os.Args[1:], " "))
	fmt.Printf("Loaded %v aliases\n", count)
	match := wtfa.FindMatch(history, aliases)
	if match != nil {
		fmt.Printf("You might try this shortcut: %v\n", match.Definition)
	} else {
		fmt.Printf("No alias found for %v\n", history.Full)
	}
}
