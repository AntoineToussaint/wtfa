package wtfa

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

type SHELL int

const (
	ZSH SHELL = iota
	BASH
)

type Stats = map[string]int

func getHistoryFile() (SHELL, string) {
	shellPath := os.Getenv("SHELL")
	fmt.Println(shellPath)
	var shell SHELL
	if strings.Contains(shellPath, "zsh") {
		shell = ZSH
	} else if strings.Contains(shellPath, "bash") {
		shell = BASH
	} else {
		fmt.Printf("%v shell not supported yet\n", shellPath)
		syscall.Exit(1)
	}
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("can't get user home directory")
		syscall.Exit(1)
	}
	var locations []string
	switch shell {
	case ZSH:
		locations = []string{fmt.Sprintf("%v/%v", dirname, ".zsh_history"), fmt.Sprintf("%v/%v", dirname, "config/zsh/.zsh_history")}
	case BASH:
		locations = []string{fmt.Sprintf("%v/%v", dirname, ".bash_history"), fmt.Sprintf("%v/%v", dirname, "config/bash/.bash_history")}
	default:
		fmt.Printf("%v not supported yet\n", shellPath)
		syscall.Exit(1)
	}
	var location string
	for _, l := range locations {
		if _, err := os.Stat(l); err == nil {
			location = l
			break
		}
	}
	return shell, location
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func GetLastCommands(n int) ([]*Cmd, Stats) {
	shell, file := getHistoryFile()
	ex := exec.Command("tail", fmt.Sprintf("-%d", n+1), file)
	out, err := ex.Output()
	if err != nil {
		fmt.Printf("can't run history command: %v\n", err)
		syscall.Exit(1)
	}
	lasts := strings.Split(string(out), "\n")
	stats := make(Stats)
	uniques := make(map[string]bool)
	// Reverse
	var cmds []*Cmd
	var cmd *Cmd
	for _, last := range reverse(lasts) {
		switch shell {
		case ZSH:
			cmd = ParseZshHistory(last)

		case BASH:
			cmd = ParseBashHistory(last)
		default:
			fmt.Println("can't parse history: unsupported shell")
			os.Exit(1)
		}
		if cmd == nil {
			continue
		}
		// Aggregate the stats
		if _, ok := stats[cmd.Full]; !ok {
			stats[cmd.Full] = 1
		} else {
			stats[cmd.Full] += 1
		}
		if _, ok := uniques[cmd.Full]; ok {
			continue
		}
		cmds = append(cmds, cmd)
		uniques[cmd.Full] = true
	}
	return cmds, stats
}

func ParseZshHistory(s string) *Cmd {
	r := regexp.MustCompile(`:\s*\d*:\d*;(.*)`)
	matches := r.FindStringSubmatch(s)
	if len(matches) < 2 {
		return nil
	}
	return ParseCommand(matches[1])
}

func ParseBashHistory(s string) *Cmd {
	r := regexp.MustCompile(`(.*)`)
	matches := r.FindStringSubmatch(s)
	if len(matches) < 2 {
		return nil
	}
	return ParseCommand(matches[1])
}
