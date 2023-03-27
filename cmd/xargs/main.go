package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mokan-r/go-cli-tools/internal/utils"
)

func root(args []string) (string, error) {
	if len(args) < 1 {
		return ``, errors.New("myXargs: you must pass a sub-command")
	}

	cmd := utils.NewCommand()
	return cmd.Run()
}

func main() {
	if result, err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Print(result)
	}
}
