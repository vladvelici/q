package main

import (
	"fmt"
	"os"

	"github.com/vladvelici/q/config"
)

func listAction(args []string) {
	fmt.Println("list action")
	fmt.Println("args are:", args)
}

func main() {

	cfgReader := config.NewReader("q")

	actionHandlers := map[string]func([]string){
		// List jobs action variants
		"list": listAction,
		"ls":   listAction,
		"ps":   listAction,

		// Add action variants
		"a":   addAction,
		"add": addAction,
		"+":   addAction,

		// Help options
		"help":   helpAction,
		"-h":     helpAction,
		"--help": helpAction,
	}

	if err := cfgReader.Flagset.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing global config. %s", err.Error())
		helpAction([]string{})
		os.Exit(2)
	}

	args := cfgReader.Flagset.Args()

	fmt.Println("Parsed args", args)

	if len(args) == 0 {
		helpAction([]string{})
		os.Exit(2)
	}

	actionStr := args[0]

	if action, ok := actionHandlers[actionStr]; ok {
		action(args[1:])
	} else {
		fmt.Fprintf(os.Stderr, "Invalid action %s\n", actionStr)
		helpAction(args[1:])
		os.Exit(2)
	}

}
