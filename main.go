package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dnlkrlv/anypy/pyenv"
)

func printAvailableVersions() {
	availableVersions := pyenv.GetAvailableVersions()
	for _, v := range availableVersions {
		fmt.Printf("- %s\n", v)
	}
}

func extractArgs() (version string, command []string) {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	version = args[0]
	command = args[1:]

	return
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: anypy [version] [command]")
		fmt.Println("Options:")
		fmt.Println("  -l, --list    List available Python versions")
		fmt.Println("  -h, --help    Show this help message")
	}

	var list bool
	flag.BoolVar(&list, "l", false, "List available Python versions")
	flag.BoolVar(&list, "list", false, "List available Python versions")

	flag.Parse()
	if list {
		printAvailableVersions()
		return
	}

	version, command := extractArgs()
	pyenv.SetPythonChain(version)

	err := pyenv.ExecuteCommand(command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
