package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Stack[T any] struct{
	items[]T
}

func (s *Stack[T]) Push(items T){
	s.items = append(s.items, items)
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}
var s Stack[string]
func main() {

	s.Push("exit")
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
			status := executeCommands(args)
			if status{
				s.Push("mac")
			}
		} else {
			status := handleBuiltin(args)
			if status{
				s.Push(line)
			}
		}
	}
}

func handleBuiltin(args []string) bool {
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "cd: expected argument")
			return false
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd:", err)
			return false
		}else{
			return true
		}
	case "exit":
		for i := 0; i < len(s.items); i++ {
			item := s.Pop()
			fmt.Printf("%s",item)
		}
		os.Exit(0)
		return true
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "exec:", err)
			return false
		}else{
			return true
		}
	}
}

func executeCommands(args []string) bool {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "mac: expected directory name")
		return false
	}
	err := os.Mkdir(args[1], 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Custom command:", err)
		return false
	}
	err = os.Chdir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "chdir:", err)
		return false
	}
	return true
}
