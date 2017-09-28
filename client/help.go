package main

import "fmt"

func helpAction(args []string) {
	helpText := `Usage: q [global-options] <action> [options]

Actions available:

- list, ls or ps    List all jobs
- add or a          Add a job
- rm                Remove a job
- help              Show this message

Use the --help flag to see the help message for a particular action.

Global options are:

-address	default 127.0.0.1:3377, the q server address.
			Also read from Q_ADDRESS environment variable, if set.

Priority is (highest to lowest): command line argument, environment variable,
default value.
`

	fmt.Println(helpText)
}
