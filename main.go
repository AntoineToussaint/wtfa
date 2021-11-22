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
	flag.IntVar(&n, "n", 100, "look for previous history, default 1")
}

func getAliases() []string {
	var aliases []string
	fileInfo, _ := os.Stdin.Stat()
	if fileInfo.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			aliases = append(aliases, scanner.Text())
		}
	}
	if len(aliases) == 0 {
		fmt.Println("wtfa requires to pipe in the alias: alias | wtfa")
		os.Exit(1)
	}
	return aliases
}

func main() {
	fmt.Println(os.Getenv("SHELL"))
	lasts := wtfa.GetLastCommands(n)
	// Pass aliases as argument
	aliases, count := wtfa.ParseAliases(getAliases())
	fmt.Printf("Loaded %v aliases!\n", count)
	matches := wtfa.FindMatches(lasts, aliases)
	blue := color.New(color.FgCyan, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	if len(matches) > 0 {
		fmt.Println("✌ We found some useful shortcuts!")
		for _, match := range matches {
			fmt.Printf("You typed: ")
			blue.Printf("%v\n", match.Cmd.Full)
			fmt.Printf("You should look at these aliases:\n")
			for _, alias := range match.Aliases {
				red.Printf("○ %v\n", alias.Definition)
			}
		}
	} else {
		fmt.Println("Looked at history")
		for _, last := range lasts {
			fmt.Println(last.Full)
		}
		fmt.Println("No alias found")
	}
}
