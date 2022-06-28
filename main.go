package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/AntoineToussaint/wtfa/wtfa"
	"github.com/fatih/color"
	"os"
)

var n int

func init() {
	flag.IntVar(&n, "n", 100, "look for previous history, default 100")
}

func getAliases() []string {
	var aliases []string
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("wtfa can't read Stdin")
		os.Exit(1)
	}
	if fileInfo.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			aliases = append(aliases, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("wtfa can't read Stdin: %v", err)
			os.Exit(1)
		}
	}
	if len(aliases) == 0 {
		fmt.Println("wtfa requires to pipe in the alias: alias | wtfa")
		fmt.Println("you can check that you have aliases available by typing alias")
		os.Exit(1)
	}
	return aliases
}

func main() {
	lasts, stats := wtfa.GetLastCommands(n)
	// Pass aliases as argument
	aliases, count := wtfa.ParseAliases(getAliases())
	fmt.Printf("Loaded %v aliases! Analyzing %v past commands.\n", count, n)
	matches, unknowns := wtfa.FindMatches(lasts, aliases)
	bold := color.New(color.FgWhite, color.Bold)
	blue := color.New(color.FgCyan, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	if len(matches) > 0 {
		_, _ = bold.Printf("♡ We found some useful existing shortcuts!\n")
		for _, match := range matches {
			fmt.Printf("Because you typed: ")
			_, _ = blue.Printf("%v\n", match.Cmd.Full)
			for _, alias := range match.Aliases {
				_, _ = red.Printf("○ %v\n", alias.Definition)
			}
		}
	}
	possibles := wtfa.Analyze(unknowns, stats, aliases)
	if len(possibles) > 0 {
		_, _ = bold.Printf("♥ You may want to create these shortcuts\n")
		for _, prop := range possibles {
			_, _ = blue.Printf("alias %v='%v'\n", prop.Shortcut, prop.Full)
		}
	}
}
