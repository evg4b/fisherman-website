package main

import (
	"fisherman/infrastructure/io"
	"fisherman/infrastructure/reporter"
	"fisherman/runner"
	"fmt"
	"os"
	"os/user"
)

func main() {
	fileAccessor := io.NewFileAccessor()
	usr, err := user.Current()
	handleError(err)
	rpt := &reporter.ConsoleReporter{}
	r := runner.NewRunner(fileAccessor, usr, rpt)
	handleError(r.Run(os.Args))
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
