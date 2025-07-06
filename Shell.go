package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	shellLoop()
}

func shellLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		cwd, _ := os.Getwd()
		fmt.Printf("msh%s> ", cwd)

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		args := strings.Fields(line)

		if args[0] == "mac" {
			executeCommands(args)
		} else {
			handleBuiltin(args)
		}
	}
}

func handleBuiltin(args []string) {
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "cd: expected argument")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd:", err)
		}
	case "exit":
		os.Exit(0)
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "exec:", err)
		}
	}
}

func executeCommands(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "mac: expected directory name")
		return
	}
	err := os.Mkdir(args[1], 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Custom command:", err)
		return
	}
	err = os.Chdir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "chdir:", err)
	}
}
