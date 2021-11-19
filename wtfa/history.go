package wtfa

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
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

func NewCmd(exec string, args []string) Cmd {
	cmd := Cmd{Exec: exec}
	for _, arg := range args {
		if arg != "" {
			cmd.Args = append(cmd.Args, arg)
		}
	}
	sort.Strings(cmd.Args)
	return cmd
}

func GetLastCommand() Cmd {
	shell, file := getHistoryFile()
	cmd := exec.Command("tail", "-2", file)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("can't run history command: %v\n", err)
		syscall.Exit(1)
	}
	last := strings.Split(string(out), "\n")[0]
	switch shell {
	case ZSH:
		return ParseZshHistory(last)
	default:
		panic("unknown shell")
	}
}

func ParseZshHistory(s string) Cmd {
	r := regexp.MustCompile(`:\s*\d*:\d*;(\w*)\s*(.*)`)
	matches := r.FindStringSubmatch(s)
	ex := matches[1]
	var args []string
	if len(matches) == 3 {
		args = strings.Split(matches[2], " ")
	}
	return NewCmd(ex, args)
}
