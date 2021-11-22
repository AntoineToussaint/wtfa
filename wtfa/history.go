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
)

func getHistoryFile() (SHELL, string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("can't get user home directory")
		syscall.Exit(1)
	}
	// TODO check if file exists, do bash...
	return ZSH, fmt.Sprintf("%v/%v", dirname, ".zsh_history")
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func GetLastCommands(n int) []*Cmd {
	shell, file := getHistoryFile()
	cmd := exec.Command("tail", fmt.Sprintf("-%d", n+1), file)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("can't run history command: %v\n", err)
		syscall.Exit(1)
	}
	lasts := strings.Split(string(out), "\n")
	uniques := make(map[string]bool)
	// Reverse
	var cmds []*Cmd
	for _, last := range reverse(lasts) {
		switch shell {
		case ZSH:
			cmd := ParseZshHistory(last)
			if cmd == nil {
				continue
			}
			if _, ok := uniques[cmd.Full]; ok {
				continue
			}
			cmds = append(cmds, cmd)
			uniques[cmd.Full] = true
		default:
			fmt.Println("shell is not supported")
		}
	}
	return cmds
}

func ParseZshHistory(s string) *Cmd {
	r := regexp.MustCompile(`:\s*\d*:\d*;(.*)`)
	matches := r.FindStringSubmatch(s)
	if len(matches) < 2 {
		return nil
	}
	return ParseCommand(matches[1])
}
