package utils

import (
	"bufio"
	"os"
	"os/exec"
)

type Runner interface {
	Run() (string, error)
	Name() string
}

type Command struct {
	command string
	args    []string
}

func (g *Command) Name() string {
	return g.command
}

func (g *Command) Run() (string, error) {
	out, err := exec.Command(g.command, g.args...).Output()
	if err == nil {
		return string(out), nil
	}

	return ``, nil
}

func ReadStdIn(command *Command) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command.args = append(command.args, scanner.Text())
	}
}

func NewCommand() *Command {
	command := new(Command)
	command.command = os.Args[1]
	command.args = append(command.args, os.Args[2:]...)

	ReadStdIn(command)

	return command
}
