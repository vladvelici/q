package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vladvelici/q/comms"
)

// getCommand checks if a command is valid, tidies it up, and returns the result.
// returns err if the path is a directory, not executable, not found in either
// $PATH or by following the path in the command...
//
// Always returns an absolute path on success.
func getCommand(command string) (string, error) {
	cmd, err := exec.LookPath(command)

	if err != nil {
		return "", err
	}

	if !filepath.IsAbs(cmd) {
		cmd, err = filepath.Abs(cmd)
	}

	return cmd, err
}

// addAction is the client method that collects the command to run and required
// environment parameters to be able to later run the job.
//
// Because of this environment encapsulation the commands can be run, for instance,
// even if they are issued inside a python virtualenv or other things that setup
// any environment variables. The job is run in exactly the same environment as
// it was queued.
func addAction(args []string) {

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "q add: No command to enqueue.")
		os.Exit(3)
	}

	command, err := getCommand(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "q add: %v\n", err)
		os.Exit(4)
	}

	env := os.Environ()

	workdir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "q add: Cannot resolve current workdir: %s\n", err.Error())
		os.Exit(4)
	}

	request := &comms.AddJobRequest{
		Script:      command,
		Workdir:     workdir,
		Args:        args[1:],
		Environment: env,
	}

	fmt.Printf("Request:\n\n%#v\n", *request)

}
