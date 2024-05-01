package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	pS "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type pipe struct {
	cmd  string
	args []string
}

func main() {
	printWorkDir()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "|")
		pipeline := []pipe{}
		for _, cmd := range fields {
			args := strings.Fields(cmd)
			pipe := pipe{cmd: args[0], args: args[1:]}
			pipeline = append(pipeline, pipe)
		}
		command(pipeline)
	}
}

func command(pipeline []pipe) {
	for _, pipe := range pipeline {
		switch pipe.cmd {
		case "quit":
			os.Exit(1)
		case "cd":
			changeDir(pipe.args)
		case "pwd":
			printWorkDir()
		case "echo":
			echo(pipe.args)
		case "kill":
			kill(pipe.args)
		case "ps":
			ps()
		default:
			execute(pipe.args)
		}
	}
}

func changeDir(args []string) {
	newDir := strings.Join(args, "")
	err := os.Chdir(newDir)
	if err != nil {
		fmt.Println(err)
	}
	printWorkDir()
}

func printWorkDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func kill(args []string) {
	pid, err := strconv.Atoi(strings.Join(args, ""))
	if err != nil {
		fmt.Println(err)
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Println(err)
	}
}

func ps() {
	processes, err := pS.Processes()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range processes {
		fmt.Printf("process name: %v	Pid: %v\n", v.Executable(), v.Pid())
	}
}

func execute(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(output) > 0 {
		fmt.Println(output)
	}
}
